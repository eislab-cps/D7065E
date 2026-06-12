# Data Engineering for Cyber-Physical Systems

Standalone notes for the third chapter of D7065E.

---

## Part 1 — Why a Whole Chapter About Data

<figure class="diagram">
<img src="figures/course-notes3-fig01.png" alt="Why data engineering matters">
<figcaption><em>A modern building generates millions of readings a day. Data engineering is the discipline that turns that flood into a trustworthy substrate an AI can act on.</em></figcaption>
</figure>

Before any AI agent can be smart, it needs data. Lots of it, well-organised, fresh enough to be useful, and clean enough to trust. The plumbing that gets raw sensor readings from the building into the agent's hands is called **data engineering**, and it is just as important as the agent itself. A brilliant model trained on bad data will produce bad predictions. A simple model trained on excellent data often beats a sophisticated model trained on garbage.

This is not just folk wisdom. In production machine-learning systems the learning algorithm itself is a small box at the centre of a vastly larger system of data collection, verification, feature extraction, and serving infrastructure — and that surrounding mass is where *hidden technical debt* accumulates ([Sculley et al., 2015](#sculley2015)). Data dependencies cost more than code dependencies precisely because they are harder to see: when an upstream sensor is recalibrated, the meaning of every feature derived from it changes silently, and no compiler raises an error. Much of this chapter can be read as a catalogue of techniques for keeping that debt visible and regularly paid down.

A useful analogy: think of the AI agent as a chef and the data pipeline as the kitchen staff who source ingredients, wash them, chop them, and lay them out on the counter. The chef gets the credit when the dish is good, but the dish was never going to be good if the ingredients arrived rotten, in the wrong portions, an hour late.

This chapter is the kitchen.

---

## Part 2 — The Three Pressures: Volume, Velocity, Variety

<figure class="diagram">
<img src="figures/course-notes3-fig02.png" alt="The three V's — volume, velocity, variety">
<figcaption><em>Volume, velocity and variety are the three pressures every data system has to absorb. They shape every choice that follows.</em></figcaption>
</figure>

A modest commercial building with 100 sensors, each reporting every 5 seconds, produces about 1.7 million readings per day. Multiply that across a year and the system holds half a billion data points. Add video cameras, vibration sensors at industrial frequencies, and access events, and the number reaches hundreds of millions of records per month.

The volume itself is not the hard part. The hard part is that data engineering for a cyber-physical system has to satisfy three demands at once, and the demands pull in different directions.

### Volume

How much data, total, accumulated over time. A weather station for one room is small. A building with thousands of sensors over several years is large. The architecture must store this without running out of disk space and without slowing down to a crawl.

An analogy: a single garden hose versus a fire hose versus a river. The amount of water you can collect determines whether you need a bucket, a tank, or a reservoir.

### Velocity

How fast new data arrives, and how fast decisions need to be made on it. A safety system that detects fire must process incoming readings in milliseconds. A weekly energy report can take its time. Both are valid, but they need different infrastructure.

Picture this: drinking a glass of water versus drinking from a fire hose. Same liquid, very different consequences.

### Variety

How many different shapes the data takes. Numeric readings, boolean states (door open/closed), strings (event types), images, time series. A pipeline that handles only one shape is brittle.

A useful image: a smoothie maker that takes apples, ice, yoghurt, and spinach — all of which need to go through different processing before they can be combined.

The fundamental tension: real-time decisions need data within seconds, machine-learning training needs months of clean labelled data, and long-term analytics needs years of queryable history. These three goals require different storage systems, different processing patterns, and different data representations. A well-designed pipeline serves all three without compromise.

Historically the industry answered this tension by building separate systems — operational databases for the fast path, data warehouses for curated analytics, data lakes for cheap raw storage — and copying data between them. Each copy step adds delay, cost, and a fresh opportunity for the copies to disagree, which is exactly the critique that motivates the lakehouse architecture discussed in Part 6 ([Armbrust et al., 2021](#armbrust2021)). Keep the three pressures in mind throughout the chapter: nearly every design decision that follows is a deliberate trade-off among them, and the honest engineering question is never "which option is best?" but "which pressure am I choosing to absorb, and at what cost to the other two?"

---

## Part 3 — Getting Sensor Data Into the System (Ingestion Patterns)

<figure class="diagram">
<img src="figures/course-notes3-fig03.png" alt="Ingestion patterns">
<figcaption><em>A broker in the middle decouples sensor producers from the consumers downstream. New consumers can subscribe without anyone changing the sensors.</em></figcaption>
</figure>

The first step in any pipeline is **ingestion**: getting data from sensors into a durable store. Three patterns dominate, each fitting a different scale.

<figure class="diagram">
<img src="figures/course-notes3-fig09.png" alt="Three ingestion patterns">
<figcaption><em>Direct REST, broker-buffered, and batch upload — the same readings, three very different sets of failure modes.</em></figcaption>
</figure>

### Pattern A: Direct REST POST

Each sensor sends every reading directly to an HTTP endpoint that writes it into a database immediately.

Easy to set up. Works fine at low volume. Breaks at scale: 100 sensors at 10 Hz means a thousand HTTP requests per second, which overwhelms most databases. There is also no buffer — if the database is slow or temporarily down, readings are lost.

Analogy: handing every letter directly to the postman, who must be standing on your doorstep every time you want to send one. Fine if you write one letter a week. Disastrous if you run a mail-order business.

### Pattern B: Message queue as buffer

Each sensor publishes to a message broker. A separate consumer process reads from the broker and writes into the database.

The broker acts as a shock absorber. If the database is briefly slow, messages accumulate in the broker, and the consumer catches up later. The sensor doesn't care whether the database is up.

Two broker families dominate in practice, and they embody different design philosophies. MQTT is a lightweight publish/subscribe protocol designed for constrained devices and unreliable networks; the broker holds messages only transiently, and its three quality-of-service levels (at-most-once, at-least-once, and an exactly-once handshake) let each sensor choose its own point on the reliability-versus-overhead curve ([OASIS, 2019](#oasis2019)). Apache Kafka takes the opposite stance: the broker is a durable, partitioned, append-only commit log that retains messages on disk for days, and consumers track their own read offsets so they can rewind and replay history at will ([Kreps et al., 2011](#kreps2011)). The log abstraction is what lets Kafka double as a short-term system of record — a crashed consumer simply resumes from its last committed offset, and a brand-new consumer can bootstrap itself from everything the log still retains. The price of decoupling, in either family, is operational: the broker becomes a critical component that must itself be monitored and, in production, replicated, and at-least-once delivery shifts the burden of deduplication onto the consumer (Part 11 returns to this).

This is the recommended pattern for the course: MQTT (using Mosquitto as the broker) plus a Python consumer that writes into TimescaleDB or DuckDB.

Picture this: the postman is replaced by a mailbox in the village square. You drop your letter in any time, the postman collects them on a schedule, the receiving end opens its mailbox when ready. Three actors, each independent.

### Pattern C: Batch file upload

The sensor accumulates readings in memory and writes them to a file (typically a Parquet file) once a minute or once an hour. A separate process reads those files into a database.

Highest throughput, lowest network overhead. Bad for real-time decisions because of the latency, excellent for historical training data.

An analogy: instead of sending one letter at a time, you wait until end-of-day and send a single envelope containing all your letters from that day. Cheap and efficient when nothing is urgent. Hopeless when something is.

Most real systems use **two patterns at once**: the broker for the real-time path, batch files for the historical path. The same reading is written into both places, and each place serves a different audience.

---

## Part 4 — Stream Processing: Working on Data as It Flows

<figure class="diagram">
<img src="figures/course-notes3-fig04.png" alt="Stream processing — rolling window">
<figcaption><em>Stream processing keeps a sliding window of recent readings in memory and produces answers continuously, with millisecond latency.</em></figcaption>
</figure>

Once data is flowing in, the next question is how to process it. Two main modes exist: streaming (processing each reading as it arrives) and batch (processing a large pile of stored data at once). This part covers streaming; the next covers batch.

A **stream** is an unbounded, time-ordered sequence of events. New events keep arriving forever. Stream processors apply operations to streams and produce either new streams or side effects (database writes, alerts, dashboard updates).

Think of it this way: a river flowing past a water-quality station. The station measures every drop as it goes past — there is no way to stop the river and inspect it all at once.

### Windowed aggregation

A single sensor reading is rarely useful by itself. It is noisy, missing context, and easily misleading. The fix is to compute statistics over a **window** of recent readings.

Three common window types, following the taxonomy given its canonical treatment in the Dataflow Model ([Akidau et al., 2015](#akidau2015)):

<figure class="diagram">
<img src="figures/course-notes3-fig10.png" alt="Three kinds of windows over the same stream">
<figcaption><em>Tumbling windows are fixed buckets, sliding windows overlap and update continuously, session windows grow until the events stop.</em></figcaption>
</figure>

**Tumbling windows** are non-overlapping, fixed-size buckets. Each reading falls into exactly one window. Useful for periodic reports: a 5-minute average that resets every 5 minutes.

**Sliding windows** are overlapping; each new reading falls into multiple windows. Useful for continuous tracking: "the rolling 5-minute average, updated every second."

**Session windows** are dynamic. They grow as long as events keep arriving and close when there's a gap. Useful for grouping bursts of activity: "all access events for one person in one trip through the building."

A natural image: imagine watching a movie versus watching the news.
- Tumbling windows are like episodes — each one starts at a fixed time and runs a fixed length.
- Sliding windows are like a live ticker that always shows "the last 5 minutes."
- Session windows are like a phone call — they start when you pick up, end when you hang up, and have no fixed length.

### Event time, processing time, and out-of-order data

Windowing looks simple in the figures above because the figures quietly assume that events arrive in order. Real streams do not behave. Every reading in a CPS carries two distinct notions of time: **event time**, when the measurement was actually taken (stamped at the sensor), and **processing time**, when the stream processor finally gets around to handling it. The two diverge whenever the network hiccups, a device buffers readings while offline, or a sensor's clock drifts (Part 11 returns to clocks). The unavoidable consequence is **out-of-order data**: a reading taken at 13:42:00 can arrive after one taken at 13:42:05, and a "complete" 5-minute window may receive a straggler long after it supposedly closed.

The canonical treatment of this problem is the Dataflow Model, distilled from Google's experience with massive unbounded streams ([Akidau et al., 2015](#akidau2015)). The core argument: for unbounded, out-of-order data, completeness can never be guaranteed — there might always be one more late event in flight — so a system must make the trade-off between correctness, latency, and cost explicit rather than pretend it away. The central mechanism is the **watermark**: a moving estimate of event-time progress that says "we believe all readings up to 13:42:00 have now arrived." When the watermark passes the end of a window, the window fires. A conservative watermark waits longer, catches more stragglers, and delays every answer; an aggressive one answers quickly but must either discard late readings or retract and amend results it has already emitted. No single setting is right for all consumers: a sprinkler decision wants low latency and can tolerate a slightly incomplete window, while a monthly energy report wants completeness and can wait. Apache Flink implements this model directly — event-time windows, watermarks, and consistent state snapshots in one engine that treats batch processing as a special case of streaming ([Carbone et al., 2015](#carbone2015)).

A familiar image: election night. The running totals shown at 22:00 are processing-time results — correct for the ballots counted so far. The official result declared days later is the event-time result, after the postal votes (the stragglers) have arrived. News channels do not refuse to show numbers until every ballot is in; they show provisional totals and revise them. Watermarks let a stream processor do the same thing, deliberately and with a tunable policy.

For the course-scale system the pragmatic policy is: window on event time using the sensor's timestamp, fall back to the server's arrival timestamp when device clocks are suspect, and configure a small allowed lateness instead of chasing perfect completeness.

### Complex event processing

Some patterns need to span multiple events over time. A fire detection rule might say: smoke above 0.7 for ten consecutive readings, combined with temperature rising at more than 2°C per minute, combined with a door opening in the same zone recently. No single reading triggers the rule; only the pattern across many readings does.

This is **complex event processing**, abbreviated CEP. Apache Flink has a CEP library that lets you describe these temporal patterns declaratively. For smaller systems, a stateful Python process maintaining a sliding window in memory is sufficient.

Analogy: a single yawn doesn't mean someone is bored. A yawn, followed by checking their phone, followed by looking at the exit, repeated three times in five minutes — that's a pattern. CEP is the skill of recognising the pattern, not the individual symptoms.

### Tools for stream processing

The right tool depends on scale.

**Apache Flink** is a distributed stream processing engine ([Carbone et al., 2015](#carbone2015)). Its distinguishing claim is unifying stream and batch processing in a single runtime, with event-time windowing, watermarks, and exactly-once state guarantees built in. Production-grade and feature-rich, but heavy to operate. Right when handling millions of events per second.

**Kafka Streams** is a stream processing library built directly on top of Apache Kafka's durable log ([Kreps et al., 2011](#kreps2011)). If Kafka is already in the architecture, Kafka Streams is the natural choice.

**Redis Streams** is a lightweight stream storage feature inside Redis. Built-in consumer groups, in-memory speed, easy to run.

**Plain Python with asyncio** is sufficient for the course. A simple consumer that reads from MQTT, computes windowed aggregations in memory, and writes to a database is entirely fine for one building.

---

## Part 5 — Batch Processing: Working on Data After It Settles

<figure class="diagram">
<img src="figures/course-notes3-fig05.png" alt="Batch processing — accumulate, then process">
<figcaption><em>Batch processing waits for data to settle, then processes large windows at once. Cheaper, simpler, but the freshest answer is from yesterday.</em></figcaption>
</figure>

Batch processing operates on large volumes of stored data all at once. A daily job that prepares training data, a weekly report that summarises energy use, a monthly anomaly analysis. Batch jobs run on a schedule (nightly, weekly) or on demand.

A useful image: streaming is fishing with a line — one fish at a time, you react as each one bites. Batch is fishing with a net — you collect a lot at once and process them together.

The streaming/batch split is a genuine trade-off, not a ranking. Streaming buys latency and pays for it with complexity: state must be kept across events, failures must be recovered mid-stream, and out-of-order data must be handled (Part 4). Batch buys throughput and simplicity — the data is complete and at rest, so a job can simply be re-run from scratch whenever its logic changes — but its freshest answer is hours old. Many architectures therefore run both modes over the same data, and modern engines increasingly blur the line from opposite directions: Flink treats a batch as a bounded stream ([Carbone et al., 2015](#carbone2015)), while Spark historically approached from the other side, treating a stream as a sequence of small batches.

For building control, the critical batch job is **training data preparation**: pulling ninety days of sensor readings out of the data lake, computing the features the ML model needs (rolling averages, occupancy patterns, time-of-day encodings), and producing a clean Parquet or CSV file for training. This job may take minutes to hours, which is fine because real-time isn't required.

### Tools for batch processing

**DuckDB** is the recommended tool for the course. It is an in-process SQL engine that queries Parquet files directly without a server. Fast, zero-configuration, full SQL. It treats a folder of Parquet files like a database table. The design point is deliberate: instead of a client–server database that data must be shipped to and from, DuckDB runs *inside* the analyst's process — the same way SQLite does for transactional workloads — which eliminates connection management and data-transfer overhead for analytical work ([Raasveldt & Mühleisen, 2019](#raasveldt2019)).

```sql
-- DuckDB reading directly from a folder of Parquet files
SELECT room, AVG(value) AS avg_temp
FROM read_parquet('data/bronze/temperature/*.parquet')
WHERE ts > now() - INTERVAL '24 hours'
GROUP BY room;
```

**pandas** is a Python DataFrame library ([McKinney, 2010](#mckinney2010)). Flexible, interactive, great for prototyping. Limited to a single machine and single thread, so it doesn't scale to industrial volumes, but for a building it is more than enough.

**Apache Spark** is a distributed batch processing engine built around resilient distributed datasets — immutable, partitioned collections that record the lineage of operations that produced them, so a lost partition can be recomputed rather than restored from replicas ([Zaharia et al., 2012](#zaharia2012)). The right tool when one machine isn't enough — multi-building analytics at industrial scale.

---

## Part 6 — The Data Lake (with the Medallion Architecture)

<figure class="diagram">
<img src="figures/course-notes3-fig06.png" alt="Medallion architecture — bronze, silver, gold">
<figcaption><em>Bronze, silver, gold. Data gets more refined and more trustworthy at each tier — and the rules that produce each tier are versioned in code.</em></figcaption>
</figure>

So far the discussion has been about *moving* data and *processing* data. The next question is where to *store* it for the long term.

### The data lake philosophy

A **data lake** is a storage system that keeps raw data in its native format, exactly as it arrived, until somebody needs it. The philosophy is "store first, structure later." Unlike a data warehouse, which requires the schema to be decided up-front, a data lake lets the schema be decided at query time.

Why does this matter? Because in a real CPS, you don't know in week one that the 5-minute variance of CO2 readings will turn out to be a useful occupancy feature. By the time you discover it, you'd be furious if you had discarded that level of detail. The data lake keeps everything, in case you need it later.

The three core principles:

1. **Store raw data.** Never transform or discard the original sensor reading. Always keep the immutable raw record.
2. **Transform on read.** Apply cleaning, feature engineering, and aggregation at query time, not at write time. This way, if you find a bug in your transformation logic, you can fix it without re-collecting the data.
3. **Schema on read.** Define the data's shape when querying, not when storing. This accommodates evolving schemas without painful migrations.

Picture this: a data lake is like a pantry that stores raw ingredients — flour, eggs, vegetables — exactly as they arrived from the grocery store. A data warehouse is like a freezer full of pre-cooked meals: faster to serve, but if you decide tomorrow that you want to make a salad instead of lasagna, you're out of luck.

The pantry has a known failure mode, though. Unstructured "store first" lakes tend to degenerate into data swamps — terabytes that nobody can find, trust, or govern — which is why enterprises long kept a separate warehouse holding curated, schema-enforced copies for the queries that mattered. The two-tier lake-plus-warehouse architecture is itself the deeper problem: every important dataset exists twice, the warehouse copy is perpetually stale, and each copy step is an opportunity for the two to disagree ([Armbrust et al., 2021](#armbrust2021)). The proposed convergence, the **lakehouse**, keeps the data in open formats such as Parquet on cheap object storage and adds a transactional metadata layer on top, so that warehouse-grade management — enforced schemas, ACID updates, versioned history, competitive query performance — runs directly against the lake ([Armbrust et al., 2021](#armbrust2021)). The medallion architecture below is best understood as the organisational discipline that makes a lake behave like a lakehouse: the raw zone stays cheap and complete, while the curated tiers supply the trust a warehouse used to provide.

### The medallion architecture

Storing everything raw is helpful, but querying everything raw every time is expensive. The compromise is the **medallion architecture**, popularised by Databricks: organise the data lake into layers, each transformed a bit more than the previous one (see the figure above). A fourth zone, the **model zone**, sits after gold.

**Bronze** is the raw zone. Each reading is stored exactly as it came off the wire, timestamped on arrival, and never modified. If a bug is found in a downstream pipeline, reprocessing from bronze is always possible.

**Silver** is the cleaned zone. Validated data (out-of-range values flagged), deduplicated (retransmissions removed), unified schema across sensors, missing values handled or marked. Silver is what most queries hit.

**Gold** is the features zone. Business-ready features for ML: rolling statistics, derived signals, time-of-day encodings, cross-sensor relationships. This is the input to model training.

**Model zone** holds ready-to-train datasets with labels and train/validation/test splits.

An analogy: bronze is your shopping receipts in a shoebox. Silver is the same receipts entered into a spreadsheet with consistent columns. Gold is a monthly summary showing categories, totals, and trends. Each is more useful than the last, but each was derived from the original receipts. If you discover an error in the gold report, you can always go back to the shoebox and reprocess.

---

## Part 7 — Storage Formats: How the Bytes Are Laid Out

A single number is just a number. A million numbers in a file is a layout decision. The chosen format affects how fast queries run, how much disk space is consumed, and how robust the data is to schema changes.

### CSV (comma-separated values)

The simplest format. One row per record, fields separated by commas, first row often contains column names.

```
ts,sensor_id,room,value
2026-04-27T13:42:00,smoke-A2306,A2306,0.82
2026-04-27T13:42:05,smoke-A2306,A2306,0.85
```

Universally readable, human-friendly, supported everywhere. But slow for large data, no schema enforcement, and stored as text (which means a 0.82 takes four bytes, not the eight bytes a double-precision float would take in binary).

Good for small exports and manual inspection. Bad for production pipelines.

### Parquet

The standard format for data lakes. Columnar binary, compressed, schema-enforced.

The key insight is **columnar storage**. CSV stores data row by row; Parquet stores it column by column:

<figure class="diagram">
<img src="figures/course-notes3-fig11.png" alt="Row storage vs columnar storage">
<figcaption><em>Row-by-row versus column-by-column layout of the same readings. The column layout is what makes Parquet small and fast.</em></figcaption>
</figure>

This matters for two reasons. First, queries that need only the `value` column don't have to read the rest. Second, columnar data compresses much better than row-based data, because consecutive values are similar (e.g., the `room` column has just a few distinct strings repeated millions of times).

Parquet files are typically 4 to 10 times smaller than equivalent CSV files, and queries that touch only a few columns run an order of magnitude faster.

Analogy: a CSV is a stack of paper forms. Each form has multiple fields, and if you want to compute the average of one field across a thousand forms, you have to flip through every form. Parquet is a spreadsheet with one column per variable; computing the average of a column is one operation on one column.

The row/column divide runs deeper than file layout. Column-store performance cannot be had simply by partitioning a row-store vertically: the real gains come from execution techniques that the columnar layout enables — compression schemes the engine can operate on without decompressing, late materialisation (delaying the reconstruction of full rows until the final result), and processing values in vectorised batches ([Abadi et al., 2008](#abadi2008)). The trade-off points the other way for writes: assembling a single new record means touching every column, so columnar formats suit data that is written once in large batches and read many times — exactly the access pattern of the silver and gold zones — and suit transactional, update-in-place workloads poorly. This is also why the hot path in Part 9 uses a row-oriented store: a stream of single-row inserts is the columnar format's worst case.

Parquet's design descends from Dremel, Google's engine for interactive analysis of web-scale datasets ([Melnik et al., 2010](#melnik2010)). From Dremel it inherits the record-shredding-and-assembly scheme — the repetition and definition levels — that lets even nested, repeated records be stored strictly column by column and reassembled losslessly on read. The practical consequence for this course: a structured JSON reading does not have to be flattened by hand before it can benefit from columnar storage.

### JSON Lines (JSONL)

One JSON object per line.

```
{"ts":"2026-04-27T13:42:00","sensor_id":"smoke-A2306","value":0.82}
{"ts":"2026-04-27T13:42:05","sensor_id":"smoke-A2306","value":0.85}
```

Flexible schema, human-readable, easy to append. Useful for event logs and audit trails where new fields are added over time.

### ORC

Similar to Parquet, used in Hadoop ecosystems. Parquet is preferred unless an existing Hive/Spark environment requires ORC.

---

## Part 8 — A Self-Hosted Data Lake with MinIO and DuckDB

For the course, an entire data lake fits on a developer's laptop using two Docker containers.

**MinIO** is an open-source object storage system that speaks the Amazon S3 API. From any S3 client's perspective, MinIO looks identical to S3, but it runs locally. Storage organisation:

```
   bucket: building-data
   ├── bronze/
   │   ├── temperature/
   │   │   ├── date=2026-04-27/readings.parquet
   │   │   └── date=2026-04-28/readings.parquet
   │   └── smoke/
   │       └── date=2026-04-27/readings.parquet
   ├── silver/
   └── gold/
```

**DuckDB** queries this directly using the `httpfs` extension. No ETL server, no cluster, no managed service. This is the in-process design point doing its work ([Raasveldt & Mühleisen, 2019](#raasveldt2019)): because DuckDB is a library embedded in the querying program rather than a server to be deployed, the "analytics cluster" for a single building collapses into a Python script and a bucket of Parquet files.

```sql
-- Query: peak smoke per sensor in the last 7 days
SELECT sensor_id,
       MAX(value) AS peak_smoke
FROM read_parquet('s3://building-data/bronze/smoke/date=*/readings.parquet')
WHERE ts > NOW() - INTERVAL '7 days'
GROUP BY sensor_id
ORDER BY peak_smoke DESC;
```

The recommended pattern for the course is to combine MinIO (the data lake) with TimescaleDB (the real-time store) and DuckDB (the SQL engine on top of MinIO). All three run in Docker. Together they cover the hot path (TimescaleDB), the cold path (MinIO + Parquet), and the analytical path (DuckDB).

---

## Part 9 — Time-Series Databases

<figure class="diagram">
<img src="figures/course-notes3-fig07.png" alt="Time-series database">
<figcaption><em>Time-series databases store rows indexed by time, compressed per chunk, and discarded automatically after a retention period. Far faster than a general-purpose database for this shape of data.</em></figcaption>
</figure>

A general-purpose database like PostgreSQL or MySQL is built for a typical web-application workload: lots of lookups by primary key, joins between tables, transactional updates. Sensor data has a completely different shape.

### Why time-series data is different

Three properties of sensor data that don't fit a general database:

1. **Append-heavy.** Every write is a new row. Existing rows are never updated. A general database's update machinery is unused overhead.
2. **Time-range queries dominate.** Almost every question is "what happened between time A and time B?" Joining tables by primary key is rare.
3. **Downsampling.** Two-year-old 5-second data is rarely needed at full resolution. Hourly averages would do, freeing 99% of the storage.

A **time-series database**, abbreviated TSDB, is built for exactly this shape. Data is stored in time-ordered chunks. New data appends to the latest chunk. Time-range predicates immediately prune entire chunks outside the range. Retention policies automatically downsample or delete old data.

Picture this: a general-purpose database is like a filing cabinet where every paper is filed by topic. Finding "all papers about Project X" is fast; finding "all papers from last March" requires walking through every folder. A time-series database is like a diary, where every page is dated and finding "what happened in March" is just opening to March.

### TSDB options

**InfluxDB** is a purpose-built time-series database. Its own data model (measurements, tags, fields) and its own query languages (InfluxQL, Flux). Excellent built-in tooling — the Telegraf agent for automatic sensor ingestion, native Grafana integration. Industry-standard for IoT.

**TimescaleDB** is a PostgreSQL extension. It adds time-series capabilities to standard PostgreSQL — data lives in regular Postgres tables, queries use standard SQL, every PostgreSQL driver works. **Hypertables** automatically partition data by time under the hood. **Continuous aggregates** keep time-windowed summaries materialised so common queries don't have to recompute them.

For the course, TimescaleDB is the most pragmatic choice: standard SQL, easy to integrate with any Python script, and it runs in one Docker container.

```sql
-- TimescaleDB: create a hypertable
CREATE TABLE readings (
    time      TIMESTAMPTZ NOT NULL,
    sensor_id TEXT NOT NULL,
    value     DOUBLE PRECISION NOT NULL,
    unit      TEXT
);
SELECT create_hypertable('readings', 'time');

-- Insert one reading
INSERT INTO readings (time, sensor_id, value, unit)
VALUES (NOW(), 'smoke-A2306', 0.82, 'normalised');

-- Query: 5-minute averages over the last 24 hours
SELECT time_bucket('5 minutes', time) AS bucket,
       AVG(value) AS avg_value,
       MAX(value) AS max_value
FROM readings
WHERE sensor_id = 'smoke-A2306'
  AND time > NOW() - INTERVAL '24 hours'
GROUP BY bucket
ORDER BY bucket;
```

**ClickHouse** is a column-oriented analytical database, extremely fast for read-heavy analytical queries for exactly the column-store reasons discussed in Part 7 ([Abadi et al., 2008](#abadi2008)). Used in production at Cloudflare and Uber. Appropriate when analytics push beyond what TimescaleDB can comfortably handle.

**DuckDB on Parquet** is the lightest option for historical analytics. No server, runs in-process, full SQL, reads Parquet files directly. Right for the cold path; less suited to high-frequency ingestion.

The recommended setup for the lab is **TimescaleDB for the hot (real-time) path and DuckDB/MinIO for the cold (historical) path**. Both use SQL. Both run in Docker.

---

## Part 10 — ETL: Extract, Transform, Load

Once data is stored, getting it into a useful shape requires a small pipeline of its own. The classic name is **ETL** — extract, transform, load.

### Extract

Reading data from one or more sources. The time-series database for recent data, the data lake for historical data, external services for context (a weather API for outdoor temperature, a calendar API for occupancy schedules).

### Transform

The step where raw data becomes useful. This is also the step that students underestimate the most.

Raw sensor readings — a float every five seconds — are not informative on their own. They are noisy, high-dimensional, and lack context. Feature engineering converts them into signals that carry meaning:

| Raw data | Derived feature | Why it helps |
|---|---|---|
| Temperature readings | 5-minute rolling average | Removes noise, reveals trend |
| Temperature readings | Rate of change (°C per minute) | Detects rapid heating (fire signature) |
| Smoke readings | Count of readings > 0.5 in last 5 minutes | More robust than a single reading |
| CO2 readings | Similarity to daily cycle | Detects occupancy anomalies |
| Door events | Time between events | Detects unusual access patterns |
| HVAC + temperature | Residual (actual − expected) | Detects HVAC failure |

An analogy: feature engineering is what a chef does to crude ingredients before serving. The raw potato is not edible. Wash, peel, slice, and bake it, and now you have something useful.

Feature pipelines are also where machine-learning systems accumulate technical debt fastest. The recurring failure modes have names: *pipeline jungles*, where data preparation grows into a tangle of ad-hoc joins and scraping steps that nobody can reproduce; *glue code*, the mass of throwaway scripts written to fit data into and out of general-purpose packages; and undeclared data dependencies, where a model silently consumes a signal whose producer does not even know it has a consumer ([Sculley et al., 2015](#sculley2015)). The prescription is the one this chapter keeps making: treat every transformation as versioned, tested code with explicit inputs and outputs — which is precisely the discipline that dbt (below) imposes on SQL transformations. When the smoke-detection pipeline in Part 12 materialises its named, reproducible feature computations into the gold zone, that is debt prevention, not bureaucracy.

### Temporal features

A surprising amount of building behaviour is predictable from the clock alone. Offices fill at 8 a.m. and empty at 6 p.m. on weekdays. CO2 rises and falls in a daily cycle. A model that knows the time of day can predict half the relevant variables before looking at any sensor.

But raw hour-and-minute values confuse machine-learning models. The model sees 23:59 and 00:01 as far apart, even though they're adjacent. The fix is **cyclical encoding** with sine and cosine.

```python
import numpy as np

def add_time_features(df):
    """Add cyclical time encodings."""
    t = df['timestamp']
    df['hour_sin'] = np.sin(2 * np.pi * t.dt.hour / 24)
    df['hour_cos'] = np.cos(2 * np.pi * t.dt.hour / 24)
    df['dow_sin']  = np.sin(2 * np.pi * t.dt.dayofweek / 7)
    df['dow_cos']  = np.cos(2 * np.pi * t.dt.dayofweek / 7)
    return df
```

The sine and cosine pair together place every hour on a unit circle. 23:59 and 00:01 end up close together on the circle, as they should.

An analogy: imagine writing the hours on a clock face. The model that uses the *clock face position* (which is what sin/cos gives it) understands that 11 o'clock and 1 o'clock are nearby. The model that uses *just the number* sees 11 and 1 as nine hours apart.

### Cross-sensor features

Relationships between sensors are often more informative than any single sensor. A few examples:

- Temperature difference between adjacent rooms — detects a door left open or an HVAC imbalance.
- CO2 combined with ventilation state — estimates occupancy without an occupancy sensor.
- Smoke level combined with temperature gradient — distinguishes cooking (high smoke, modest temperature rise) from a real fire (high smoke, fast temperature rise).

### Load

Writing the computed features to a feature store — a database or Parquet file ready for ML training — or feeding them directly into the model.

### ETL tools

**dbt (data build tool)** is a popular open-source tool that lets transformations be expressed as SQL queries. Dbt runs queries in the correct order, tests their outputs, and generates documentation. Free, well-supported, with a free 4-hour fundamentals course.

**pandas** is the standard Python DataFrame library ([McKinney, 2010](#mckinney2010)). Flexible and interactive, ideal for exploratory work. Single-threaded and in-memory, so it doesn't scale to industrial volumes; production pipelines are usually rewritten in SQL (using dbt and DuckDB) once the transformations are stable.

**Apache Airflow** is a workflow orchestration platform. ETL jobs are defined as directed acyclic graphs of tasks, scheduled, retried on failure, and alerted on errors. Production-grade. For a course-scale system, a Python script scheduled with cron is sufficient; Airflow is the heavy-machinery version.

---

## Part 11 — Data Quality and Observability

Data quality is the silent killer of machine-learning systems. A model trained on bad data will produce bad predictions, and the failures are often subtle. The model performs well on the training data (because the training data is also bad), then fails mysteriously in production.

For a CPS, bad data is not just an ML problem. It is a safety problem. A smoke sensor stuck at zero will silently prevent the fire-detection system from ever responding.

### Five common data-quality failures

**Missing data.** A sensor goes offline because the network drops, the power fails, or the sensor itself dies. The pipeline receives no readings for some period. Handling options:

- *Forward-fill*: use the last known value. Reasonable for slowly changing quantities like temperature. Dangerous for fast-changing ones like door state.
- *Linear interpolation*: estimate values based on readings before and after the gap. Reasonable for smooth signals.
- *Mark as missing*: insert a null marker that the ML model handles explicitly.
- *Alert on absence*: if a safety sensor has not reported in 30 seconds, raise an alarm. Don't fill silently — react.

**Duplicate data.** MQTT's at-least-once delivery means a message may be delivered twice. Without protection, every duplicated smoke reading inflates the model's view of how often anomalies happen. The fix is idempotent inserts (write only if not already present) or dedup logic in the consumer that keys on `(sensor_id, ts)`.

Analogy: imagine if every time the postman wasn't sure whether a letter was delivered, he delivered it twice. Soon your filing cabinet has two of every letter. You need a rule: if the same letter arrives twice, throw the second one away.

**Stale data.** The sensor process is alive, the network is fine, but the value never changes. A frozen sensor reads the same value indefinitely. Detection: compute the standard deviation of readings over a 5-minute window — if it's exactly zero, the sensor is probably stuck. Monitoring: track each sensor's last-updated timestamp and alert if nothing has changed for N seconds.

**Clock drift.** Each sensor has its own clock, and clocks drift. A reading from sensor A timestamped 14:32:00 and one from sensor B timestamped 14:31:58 may have happened in the opposite order from what the timestamps suggest. Mitigations: run NTP (Network Time Protocol) on every device, store both the device timestamp and the time when the server received the message, and prefer the server timestamp for ordering. NTP version 4 disciplines each device's clock against a hierarchy of reference servers and, on a typical local network, holds devices within roughly a millisecond of one another ([IETF RFC 5905, 2010](#rfc5905)) — easily good enough to order readings taken seconds apart, though not good enough for sub-millisecond event correlation. Synchronised clocks are also what makes the event-time windowing of Part 4 trustworthy: a watermark is only as honest as the timestamps beneath it.

Picture this: imagine a courtroom where every witness uses a different clock. Their statements about timing can't be compared without first synchronising the clocks.

**Schema evolution.** A new sensor type is added with an extra field. The existing pipeline doesn't know what to do with that field. Schema-on-read (typical for data lakes) handles this gracefully — old code ignores new fields. Schema-on-write (typical for relational databases) requires a migration to add the column.

### Monitoring the pipeline

A data pipeline that silently fails is worse than one that fails loudly. Production-quality pipelines instrument themselves with metrics, and a monitoring system alerts when something looks wrong.

Four metrics worth tracking:

- **Data freshness.** Age of the most recent reading for each sensor. Alert if it exceeds three times the expected interval.
- **Value range.** A temperature below -50°C or above 100°C inside a Swedish office building is impossible. Alert when readings violate physical bounds.
- **Volume.** Number of readings per minute. A sudden drop signals offline sensors or a broken consumer.
- **Error rate.** Number of readings rejected by validation. A spike suggests a schema change or a faulty sensor.

Note the division of labour: these checks catch *faulty data* with simple physical rules, while the anomaly model in Part 12 catches *anomalous reality* with statistics. The boundary is genuinely blurry — a stuck sensor and a real fire can look alike from a single channel — and the research literature on anomaly detection treats sensor faults and real events within the same formal framework ([Chandola et al., 2009](#chandola2009)). The practical consequence is an ordering rule: validate data quality *before* the anomaly model, so the model spends its statistical power on the building rather than on the plumbing.

The standard stack for this is **Prometheus** (a metrics database that scrapes endpoints periodically) plus **Grafana** (a dashboard tool that visualises the metrics and triggers alerts). Both are free, both run in Docker, both are industry standard.

---

## Part 12 — A Worked Example: Smoke Detection Data Pipeline

<figure class="diagram">
<img src="figures/course-notes3-fig08.png" alt="Smoke detection data pipeline">
<figcaption><em>Every component of the smoke-detection data pipeline: sensors at the top push readings down through ingestion, processing, modelling, and alerting — and everything is persisted to storage for replay and training.</em></figcaption>
</figure>

To make all of the above concrete, follow one use case — smoke detection — through every stage.

**Step 1: Ingestion.** Each smoke sensor (one per room) publishes a reading every 5 seconds. The reading is a JSON object:

```json
{
  "ts": "2026-04-27T13:42:00.250Z",
  "sensor_id": "smoke-A2306",
  "room": "A2306",
  "level": "level1",
  "type": "smoke",
  "unit": "fraction",
  "value": 0.04
}
```

Published to MQTT topic `sensors/level1/A2306/smoke` at QoS 1.

**Step 2: Stream consumer.** A small Python process subscribes to `sensors/#`. For every incoming message, it inserts a row into the TimescaleDB `readings` hypertable (the hot path) and appends to today's Parquet file in MinIO under `s3://building-data/bronze/smoke/date=2026-04-27/readings.parquet` (the cold path).

**Step 3: Stream processing.** A second small process keeps an in-memory 5-minute sliding window per sensor. For each new reading, it recomputes the count of readings above 0.5 in that window. If the count exceeds a threshold, it publishes to `alerts/anomaly/A2306`.

**Step 4: Anomaly model (real-time).** The safety agent subscribes to `alerts/anomaly/+`. When an alert arrives, it queries TimescaleDB for the last 60 seconds of smoke, temperature, and CO2 in that room, computes a feature vector, and runs the Isolation Forest model. If the score is above 0.8, it issues a sprinkler command.

A note on the model. Isolation Forest is an unsupervised ensemble method that isolates anomalies through random recursive partitioning: anomalous points are few and different, so they separate from the rest in fewer splits and sit closer to the roots of the trees ([Liu et al., 2008](#liu2008)). Within the standard taxonomy of anomaly-detection techniques it is an unsupervised point-anomaly detector — a sensible default when labelled examples of real fires are, thankfully, scarce ([Chandola et al., 2009](#chandola2009)). Do not implement the algorithm yourself: use the scikit-learn implementation ([Pedregosa et al., 2011](#pedregosa2011)), documented at https://scikit-learn.org/stable/modules/generated/sklearn.ensemble.IsolationForest.html. The engineering effort in this course belongs in the feature vector that goes *into* the model — the rolling statistics and cross-sensor features of Part 10 — not in re-deriving the algorithm.

**Step 5: Batch training (nightly).** A scheduled job runs at 3 a.m. each night. It reads the previous 24 hours of bronze data with DuckDB, cleans it (silver), computes ML features (gold), and writes the result to `s3://building-data/gold/smoke-features/date=2026-04-27/`. The training script reads the gold file, trains an updated anomaly model, and writes the new model to a model store. The safety agent picks up the new model on its next restart.

**Step 6: Monitoring.** A dashboard in Grafana shows per-sensor freshness, the readings-per-minute rate, the false-positive rate of the anomaly model over the last week, and the percentage of readings rejected by validation. An alert fires when any sensor's freshness exceeds 30 seconds or when the false-positive rate exceeds 5 percent.

That sequence is a complete, production-shaped data pipeline. Bronze, silver, gold, hot path, cold path, monitoring. Every step uses tools that fit on a laptop. Every step generates artefacts that survive across restarts.

---

## Part 13 — Vocabulary Reference

Every term used in this chapter, defined.

| Term | Definition |
|---|---|
| **Data engineering** | The discipline of designing, building, and maintaining the systems that move data from sources to consumers |
| **Ingestion** | The act of getting data from a producer (sensor) into a durable store |
| **Stream processing** | Operating on data as it arrives, one event at a time, without first storing it |
| **Batch processing** | Operating on large piles of stored data all at once on a schedule |
| **Window (in streaming)** | A time-bounded subset of a stream over which aggregations are computed |
| **Tumbling window** | A non-overlapping window of fixed size |
| **Sliding window** | An overlapping window of fixed size that updates continuously |
| **Session window** | A dynamic window that closes when events stop arriving for a gap |
| **Event time** | The time at which a reading was actually taken, stamped at the source |
| **Processing time** | The time at which the stream processor handles a reading |
| **Watermark** | A stream processor's moving estimate of event-time progress, used to decide when a window is complete |
| **Complex event processing (CEP)** | Detecting temporal patterns across multiple events |
| **Data lake** | A storage system that keeps raw data in its native format until it is needed |
| **Lakehouse** | An architecture that adds warehouse-grade management (schemas, transactions, governance) directly on open data-lake storage |
| **Medallion architecture** | A pattern that organises a data lake into bronze (raw), silver (cleaned), gold (features), and model zones |
| **Bronze layer** | The raw, immutable copy of incoming data |
| **Silver layer** | Cleaned, deduplicated, validated data with a unified schema |
| **Gold layer** | ML-ready features derived from silver |
| **Parquet** | A columnar binary file format, compressed and schema-enforced, used for data lakes |
| **CSV** | Comma-separated text file format; simple but slow |
| **JSON Lines (JSONL)** | One JSON object per line; flexible schema for event logs |
| **Schema on write** | The schema is enforced when data is stored (e.g., relational databases) |
| **Schema on read** | The schema is applied when data is queried (e.g., data lakes) |
| **Object storage** | A storage system that holds files (objects) in flat buckets, with HTTP API |
| **MinIO** | An open-source object storage system compatible with the Amazon S3 API |
| **Time-series database (TSDB)** | A database optimised for time-stamped data, append-heavy writes, and time-range queries |
| **InfluxDB** | A purpose-built time-series database popular in IoT |
| **TimescaleDB** | A PostgreSQL extension that adds time-series capabilities |
| **Hypertable** | A TimescaleDB construct that partitions data by time automatically |
| **DuckDB** | An in-process SQL engine that reads Parquet files directly |
| **ClickHouse** | A column-oriented analytical database for read-heavy workloads |
| **ETL (Extract, Transform, Load)** | A pipeline that reads data, reshapes it, and writes it to a downstream store |
| **Feature engineering** | The process of converting raw data into informative signals for ML |
| **Rolling average** | The average of a value over a recent window of time |
| **Cyclical encoding** | Representing periodic variables (hour, day-of-week) using sine and cosine to preserve their adjacency |
| **Feature store** | A database that holds ML-ready features ready to be consumed by training and inference |
| **Isolation Forest** | An unsupervised anomaly-detection algorithm that isolates outliers through random recursive partitioning |
| **dbt** | A tool for expressing data transformations as SQL queries with built-in testing |
| **Apache Airflow** | A workflow orchestration platform that runs ETL pipelines as scheduled DAGs |
| **NTP (Network Time Protocol)** | A protocol for synchronising clocks across networked devices |
| **Prometheus** | A metrics database that scrapes data from instrumented services |
| **Grafana** | A dashboard tool that visualises metrics and triggers alerts |
| **Idempotent insert** | An insert that produces the same result whether performed once or many times |

---

## Part 14 — Summary in Five Sentences

1. Sensor data has a shape unlike most software data: high volume, fast velocity, mixed variety, and a hard tension between real-time decisions, ML training, and long-term analytics.
2. Ingestion uses a message broker as a buffer between sensors and storage, so that producers and consumers are independent and a slow downstream system does not lose data.
3. Storage is split between a hot path (TimescaleDB for real-time queries) and a cold path (Parquet on MinIO for historical analytics), with the medallion architecture (bronze, silver, gold, model) organising the cold path.
4. Feature engineering — rolling averages, rate of change, cyclical time encodings, cross-sensor relationships — turns raw readings into signals the ML model can learn from.
5. Data quality is a safety concern, not just an analytics concern; pipelines must detect missing, duplicate, stale, and clock-drifted data, and they must instrument themselves with metrics that signal when something is going wrong.

These five ideas are the foundation for every downstream use of data in this course — for AI agents, dashboards, training runs, and audit trails.

---

## Part 15 — References

### Literature

- <a id="abadi2008"></a>Abadi, D. J., Madden, S. R., & Hachem, N. (2008). Column-stores vs. row-stores: How different are they really? In *Proceedings of the ACM SIGMOD International Conference on Management of Data*.
- <a id="akidau2015"></a>Akidau, T., Bradshaw, R., Chambers, C., Chernyak, S., Fernández-Moctezuma, R. J., Lax, R., McVeety, S., Mills, D., Perry, F., Schmidt, E., & Whittle, S. (2015). The Dataflow Model: A practical approach to balancing correctness, latency, and cost in massive-scale, unbounded, out-of-order data processing. *Proceedings of the VLDB Endowment*, 8(12).
- <a id="armbrust2021"></a>Armbrust, M., Ghodsi, A., Xin, R., & Zaharia, M. (2021). Lakehouse: A new generation of open platforms that unify data warehousing and advanced analytics. In *Proceedings of the Conference on Innovative Data Systems Research (CIDR)*.
- <a id="carbone2015"></a>Carbone, P., Katsifodimos, A., Ewen, S., Markl, V., Haridi, S., & Tzoumas, K. (2015). Apache Flink: Stream and batch processing in a single engine. *IEEE Data Engineering Bulletin*, 38(4).
- <a id="chandola2009"></a>Chandola, V., Banerjee, A., & Kumar, V. (2009). Anomaly detection: A survey. *ACM Computing Surveys*, 41(3).
- <a id="kreps2011"></a>Kreps, J., Narkhede, N., & Rao, J. (2011). Kafka: A distributed messaging system for log processing. In *Proceedings of the NetDB Workshop*.
- <a id="liu2008"></a>Liu, F. T., Ting, K. M., & Zhou, Z.-H. (2008). Isolation Forest. In *Proceedings of the IEEE International Conference on Data Mining (ICDM)*.
- <a id="mckinney2010"></a>McKinney, W. (2010). Data structures for statistical computing in Python. In *Proceedings of the 9th Python in Science Conference (SciPy)*.
- <a id="melnik2010"></a>Melnik, S., Gubarev, A., Long, J. J., Romer, G., Shivakumar, S., Tolton, M., & Vassilakis, T. (2010). Dremel: Interactive analysis of web-scale datasets. *Proceedings of the VLDB Endowment*, 3(1).
- <a id="pedregosa2011"></a>Pedregosa, F., et al. (2011). Scikit-learn: Machine learning in Python. *Journal of Machine Learning Research*, 12.
- <a id="raasveldt2019"></a>Raasveldt, M., & Mühleisen, H. (2019). DuckDB: An embeddable analytical database. In *Proceedings of the ACM SIGMOD International Conference on Management of Data* (demonstration).
- <a id="sculley2015"></a>Sculley, D., Holt, G., Golovin, D., Davydov, E., Phillips, T., Ebner, D., Chaudhary, V., Young, M., Crespo, J.-F., & Dennison, D. (2015). Hidden technical debt in machine learning systems. In *Advances in Neural Information Processing Systems (NeurIPS)*.
- <a id="zaharia2012"></a>Zaharia, M., Chowdhury, M., Das, T., Dave, A., Ma, J., McCauley, M., Franklin, M. J., Shenker, S., & Stoica, I. (2012). Resilient distributed datasets: A fault-tolerant abstraction for in-memory cluster computing. In *Proceedings of the USENIX Symposium on Networked Systems Design and Implementation (NSDI)*.

### Software, standards, and online resources

- Apache Airflow, workflow orchestration platform. https://airflow.apache.org
- Apache Parquet, columnar storage format. https://parquet.apache.org
- ClickHouse, column-oriented analytical database. https://clickhouse.com
- dbt (data build tool), SQL transformation framework. https://www.getdbt.com
- DuckDB, in-process analytical database. https://duckdb.org
- Eclipse Mosquitto, open-source MQTT broker. https://mosquitto.org
- Grafana, dashboards and alerting. https://grafana.com
- <a id="rfc5905"></a>IETF RFC 5905 (2010). Network Time Protocol Version 4: Protocol and Algorithms Specification.
- InfluxDB, time-series database. https://www.influxdata.com
- MinIO, S3-compatible object storage. https://min.io
- <a id="oasis2019"></a>OASIS (2019). MQTT Version 5.0, OASIS Standard. https://mqtt.org
- Prometheus, metrics database and monitoring system. https://prometheus.io
- scikit-learn, IsolationForest documentation. https://scikit-learn.org/stable/modules/generated/sklearn.ensemble.IsolationForest.html
- TimescaleDB, time-series extension for PostgreSQL. https://www.timescale.com
