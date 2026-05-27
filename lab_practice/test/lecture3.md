# Lecture 3 — Data Engineering for Cyber-Physical Systems

Standalone notes for the third lecture of D7065E. Read on its own or alongside `lectures/lecture-3-data-engineering.md`.

---

## Part 1 — Why a Whole Lecture About Data

<figure class="diagram"><div class="dgm-frame">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 720 280" role="img" aria-label="Why data engineering matters" class="dgm">
<text x="360" y="30" text-anchor="middle" font-size="14" font-weight="600" fill="#2a2622">From a flood of readings to something usable</text>
<!-- Many sensors emitting dots -->
<g transform="translate(40,80)">
  <g transform="translate(0,0)">
    <circle cx="0" cy="0" r="10" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.4"/>
    <circle cx="0" cy="0" r="3" fill="#8b3a1f"/>
    <path d="M -8 -10 Q -12 -14 -8 -18" fill="none" stroke="#6b6660" stroke-width="1.2"/>
    <path d="M 8 -10 Q 12 -14 8 -18" fill="none" stroke="#6b6660" stroke-width="1.2"/>
  </g><g transform="translate(0,40)">
    <circle cx="0" cy="0" r="10" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.4"/>
    <circle cx="0" cy="0" r="3" fill="#8b3a1f"/>
    <path d="M -8 -10 Q -12 -14 -8 -18" fill="none" stroke="#6b6660" stroke-width="1.2"/>
    <path d="M 8 -10 Q 12 -14 8 -18" fill="none" stroke="#6b6660" stroke-width="1.2"/>
  </g><g transform="translate(0,80)">
    <circle cx="0" cy="0" r="10" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.4"/>
    <circle cx="0" cy="0" r="3" fill="#8b3a1f"/>
    <path d="M -8 -10 Q -12 -14 -8 -18" fill="none" stroke="#6b6660" stroke-width="1.2"/>
    <path d="M 8 -10 Q 12 -14 8 -18" fill="none" stroke="#6b6660" stroke-width="1.2"/>
  </g><g transform="translate(0,120)">
    <circle cx="0" cy="0" r="10" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.4"/>
    <circle cx="0" cy="0" r="3" fill="#8b3a1f"/>
    <path d="M -8 -10 Q -12 -14 -8 -18" fill="none" stroke="#6b6660" stroke-width="1.2"/>
    <path d="M 8 -10 Q 12 -14 8 -18" fill="none" stroke="#6b6660" stroke-width="1.2"/>
  </g><g transform="translate(0,160)">
    <circle cx="0" cy="0" r="10" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.4"/>
    <circle cx="0" cy="0" r="3" fill="#8b3a1f"/>
    <path d="M -8 -10 Q -12 -14 -8 -18" fill="none" stroke="#6b6660" stroke-width="1.2"/>
    <path d="M 8 -10 Q 12 -14 8 -18" fill="none" stroke="#6b6660" stroke-width="1.2"/>
  </g>
</g>
<!-- chaotic dots flowing right -->
<circle cx="232.26884024831867" cy="217.85485277327285" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="184.02820025246598" cy="169.67505126691213" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="203.55902171535809" cy="101.41208450312848" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="157.53763187752065" cy="125.29103956464519" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="126.97615660814257" cy="190.6462483792548" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="210.7811800539409" cy="102.87000947989311" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="120.88048814096327" cy="105.07504735462852" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="90.25789528551493" cy="85.81127280101046" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="247.7550312828177" cy="172.51756548768833" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="119.49398758604417" cy="229.3076429587822" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="199.42207310845623" cy="89.94385842604763" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="169.7979951065526" cy="240.39594939563082" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="170.507317870907" cy="127.22437837664198" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="122.52100127144699" cy="182.1702997870683" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="187.40356490994031" cy="153.20561110625908" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="115.49877942305307" cy="165.70739644292183" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="154.55104701103684" cy="195.2341802939285" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="226.24026397798548" cy="96.51314092012117" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="243.44706959145026" cy="92.09509879512301" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="202.1852128777553" cy="85.49238701846333" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="219.22041818594025" cy="141.39349929785448" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="207.46537579243895" cy="205.38272594652932" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="183.27945574150425" cy="194.37400303512902" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="203.271882403615" cy="132.89357882432796" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="120.97985563182502" cy="198.1440546362665" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="171.76392327873978" cy="221.21280584632515" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="226.2411529071279" cy="213.38455522371405" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="223.2217584711932" cy="234.56271255909294" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="102.21342090942909" cy="162.85757980235968" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="244.4328410983301" cy="113.5747482459523" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="158.44221488296358" cy="179.67364561406933" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="93.08254924374593" cy="137.45632983535313" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="212.1693672628211" cy="174.49664467012582" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="152.59737676551555" cy="154.49109434353613" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="187.18623284119593" cy="239.97168415355432" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="106.37115038139024" cy="174.7229756871927" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="145.21153180984714" cy="96.71923106648248" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="196.7022780982079" cy="222.3811923514005" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="225.78676965398265" cy="108.33048751200374" r="2.2" fill="#6b6660" opacity="0.55"/><circle cx="144.50540586797158" cy="92.641118179716" r="2.2" fill="#6b6660" opacity="0.55"/>
<!-- Funnel -->
<polygon points="280,90 480,90 410,180 350,180" fill="#f3ece6" stroke="#2a2622" stroke-width="1.5"/>
<text x="380" y="140" text-anchor="middle" font-size="13" font-weight="700" fill="#8b3a1f">DATA</text>
<text x="380" y="158" text-anchor="middle" font-size="13" font-weight="700" fill="#8b3a1f">ENGINEERING</text>
<!-- Clean stream out -->
<path d="M 380 182 L 380 230" fill="none" stroke="#8b3a1f" stroke-width="2"/>
<circle cx="430" cy="200" r="2.5" fill="#8b3a1f"/><circle cx="430" cy="205" r="2.5" fill="#8b3a1f"/><circle cx="430" cy="210" r="2.5" fill="#8b3a1f"/><circle cx="430" cy="215" r="2.5" fill="#8b3a1f"/><circle cx="430" cy="220" r="2.5" fill="#8b3a1f"/>
<line x1="395" y1="200" x2="533" y2="200" stroke="#8b3a1f" stroke-width="2" stroke-linecap="round"/>
<polygon points="540,200 533,204 533,196" fill="#8b3a1f"/>
<g transform="translate(560,180)">
  <rect width="120" height="40" rx="6" fill="#e2ebde" stroke="#2a2622" stroke-width="1.5"/>
  <text x="60" y="24" text-anchor="middle" font-size="12" font-weight="600" fill="#2a2622">Trustworthy</text>
</g>
<text x="160" y="270" text-anchor="middle" font-size="10.5" font-style="italic" fill="#6b6660">millions of raw readings/day</text>
<text x="620" y="240" text-anchor="middle" font-size="10.5" font-style="italic" fill="#6b6660">substrate for AI</text>
</svg>
</div><figcaption>A modern building generates millions of readings a day. Data engineering is the discipline that turns that flood into a trustworthy substrate an AI can act on.</figcaption></figure>

Before any AI agent can be smart, it needs data. Lots of it, well-organised, fresh enough to be useful, and clean enough to trust. The plumbing that gets raw sensor readings from the building into the agent's hands is called **data engineering**, and it is just as important as the agent itself. A brilliant model trained on bad data will produce bad predictions. A simple model trained on excellent data often beats a sophisticated model trained on garbage.

A useful analogy: think of the AI agent as a chef and the data pipeline as the kitchen staff who source ingredients, wash them, chop them, and lay them out on the counter. The chef gets the credit when the dish is good, but the dish was never going to be good if the ingredients arrived rotten, in the wrong portions, an hour late.

This chapter is the kitchen.

---

## Part 2 — The Three Pressures: Volume, Velocity, Variety

<figure class="diagram"><div class="dgm-frame">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 720 320" role="img" aria-label="The three V\u2019s — volume, velocity, variety" class="dgm">
<text x="360" y="30" text-anchor="middle" font-size="14" font-weight="600" fill="#2a2622">Three pressures every data system feels</text>
  <g transform="translate(60,70)">
    <rect width="160" height="220" rx="14" fill="#f3ece6" stroke="#2a2622" stroke-width="1.5"/>
    <rect x="40" y="64" width="80" height="20" fill="#8b3a1f" opacity="0.85"/>
    <rect x="40" y="86" width="80" height="20" fill="#8b3a1f" opacity="0.7"/>
    <rect x="40" y="108" width="80" height="20" fill="#8b3a1f" opacity="0.55"/>
    <rect x="40" y="130" width="80" height="20" fill="#8b3a1f" opacity="0.4"/>
    <text x="80" y="50" text-anchor="middle" font-size="14" font-weight="700" fill="#8b3a1f">TB</text>
    <text x="80" y="190" text-anchor="middle" font-size="15" font-weight="700" fill="#2a2622">Volume</text>
    <text x="80" y="208" text-anchor="middle" font-size="11" fill="#6b6660" font-style="italic">how much?</text>
  </g>
  <g transform="translate(280,70)">
    <rect width="160" height="220" rx="14" fill="#dde7ec" stroke="#2a2622" stroke-width="1.5"/>
    <circle cx="80" cy="100" r="34" fill="none" stroke="#8b3a1f" stroke-width="2"/>
    <line x1="80" y1="100" x2="80" y2="72" stroke="#8b3a1f" stroke-width="2.5"/>
    <line x1="80" y1="100" x2="100" y2="100" stroke="#8b3a1f" stroke-width="2.5"/>
    <text x="80" y="146" text-anchor="middle" font-size="11" fill="#2a2622">100 Hz · per sensor</text>
    <text x="80" y="190" text-anchor="middle" font-size="15" font-weight="700" fill="#2a2622">Velocity</text>
    <text x="80" y="208" text-anchor="middle" font-size="11" fill="#6b6660" font-style="italic">how fast?</text>
  </g>
  <g transform="translate(500,70)">
    <rect width="160" height="220" rx="14" fill="#e2ebde" stroke="#2a2622" stroke-width="1.5"/>
    <rect x="40" y="72" width="22" height="22" rx="3" fill="#8b3a1f" opacity="0.7"/>
    <circle cx="80" cy="83" r="11" fill="#8b3a1f" opacity="0.55"/>
    <polygon points="108,72 122,94 94,94" fill="#8b3a1f" opacity="0.4"/>
    <rect x="56" y="106" width="18" height="18" rx="9" fill="#8b3a1f" opacity="0.85"/>
    <polygon points="90,108 102,108 102,120 90,120" fill="#8b3a1f" opacity="0.3"/>
    <text x="80" y="156" text-anchor="middle" font-size="10" fill="#2a2622">JSON · CSV · binary</text>
    <text x="80" y="190" text-anchor="middle" font-size="15" font-weight="700" fill="#2a2622">Variety</text>
    <text x="80" y="208" text-anchor="middle" font-size="11" fill="#6b6660" font-style="italic">how many shapes?</text>
  </g>
</svg>
</div><figcaption>Volume, velocity and variety are the three pressures every data system has to absorb. They shape every choice that follows.</figcaption></figure>

A modest commercial building with 100 sensors, each reporting every 5 seconds, produces about 1.7 million readings per day. Multiply that across a year and the system holds half a billion data points. Add video cameras, vibration sensors at industrial frequencies, and access events, and the number reaches hundreds of millions of records per month.

The volume itself is not the hard part. The hard part is that data engineering for a cyber-physical system has to satisfy three demands at once, and the demands pull in different directions.

### Volume

How much data, total, accumulated over time. A weather station for one room is small. A building with thousands of sensors over several years is large. The architecture must store this without running out of disk space and without slowing down to a crawl.

A useful image: a single garden hose versus a fire hose versus a river. The amount of water you can collect determines whether you need a bucket, a tank, or a reservoir.

### Velocity

How fast new data arrives, and how fast decisions need to be made on it. A safety system that detects fire must process incoming readings in milliseconds. A weekly energy report can take its time. Both are valid, but they need different infrastructure.

A useful image: drinking a glass of water versus drinking from a fire hose. Same liquid, very different consequences.

### Variety

How many different shapes the data takes. Numeric readings, boolean states (door open/closed), strings (event types), images, time series. A pipeline that handles only one shape is brittle.

A useful image: a smoothie maker that takes apples, ice, yoghurt, and spinach — all of which need to go through different processing before they can be combined.

The fundamental tension: real-time decisions need data within seconds, machine-learning training needs months of clean labelled data, and long-term analytics needs years of queryable history. These three goals require different storage systems, different processing patterns, and different data representations. A well-designed pipeline serves all three without compromise.

---

## Part 3 — Getting Sensor Data Into the System (Ingestion Patterns)

<figure class="diagram"><div class="dgm-frame">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 720 320" role="img" aria-label="Ingestion patterns" class="dgm">
<text x="360" y="30" text-anchor="middle" font-size="14" font-weight="600" fill="#2a2622">A broker lets producers and consumers evolve apart</text>
<!-- Sensors -->
<g transform="translate(40,70)">
  <g transform="translate(0,0)"><g transform="translate(20,16)">
    <circle cx="0" cy="0" r="10" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.4"/>
    <circle cx="0" cy="0" r="3" fill="#8b3a1f"/>
    <path d="M -8 -10 Q -12 -14 -8 -18" fill="none" stroke="#6b6660" stroke-width="1.2"/>
    <path d="M 8 -10 Q 12 -14 8 -18" fill="none" stroke="#6b6660" stroke-width="1.2"/>
  </g><text x="44" y="20" font-size="10.5" fill="#2a2622">sensor 1</text></g><g transform="translate(0,50)"><g transform="translate(20,16)">
    <circle cx="0" cy="0" r="10" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.4"/>
    <circle cx="0" cy="0" r="3" fill="#8b3a1f"/>
    <path d="M -8 -10 Q -12 -14 -8 -18" fill="none" stroke="#6b6660" stroke-width="1.2"/>
    <path d="M 8 -10 Q 12 -14 8 -18" fill="none" stroke="#6b6660" stroke-width="1.2"/>
  </g><text x="44" y="20" font-size="10.5" fill="#2a2622">sensor 2</text></g><g transform="translate(0,100)"><g transform="translate(20,16)">
    <circle cx="0" cy="0" r="10" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.4"/>
    <circle cx="0" cy="0" r="3" fill="#8b3a1f"/>
    <path d="M -8 -10 Q -12 -14 -8 -18" fill="none" stroke="#6b6660" stroke-width="1.2"/>
    <path d="M 8 -10 Q 12 -14 8 -18" fill="none" stroke="#6b6660" stroke-width="1.2"/>
  </g><text x="44" y="20" font-size="10.5" fill="#2a2622">sensor 3</text></g><g transform="translate(0,150)"><g transform="translate(20,16)">
    <circle cx="0" cy="0" r="10" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.4"/>
    <circle cx="0" cy="0" r="3" fill="#8b3a1f"/>
    <path d="M -8 -10 Q -12 -14 -8 -18" fill="none" stroke="#6b6660" stroke-width="1.2"/>
    <path d="M 8 -10 Q 12 -14 8 -18" fill="none" stroke="#6b6660" stroke-width="1.2"/>
  </g><text x="44" y="20" font-size="10.5" fill="#2a2622">sensor 4</text></g>
</g>
<!-- Arrows in -->
<line x1="140" y1="86" x2="254.04180391193208" y2="156.32577907902478" stroke="#8b3a1f" stroke-width="1.4" stroke-linecap="round"/>
<polygon points="260,160 251.94224909994622,159.73046255792073 256.1413587239179,152.92109560012884" fill="#8b3a1f"/><line x1="140" y1="136" x2="253.13593527016357" y2="158.6271870540327" stroke="#8b3a1f" stroke-width="1.4" stroke-linecap="round"/>
<polygon points="260,160 252.35147072961084,162.5495097567964 253.9203998107163,154.704864351269" fill="#8b3a1f"/><line x1="140" y1="186" x2="253.15873801981243" y2="161.48227342904065" stroke="#8b3a1f" stroke-width="1.4" stroke-linecap="round"/>
<polygon points="260,160 254.00575140783565,165.39156598914784 252.31172463178922,157.57298086893346" fill="#8b3a1f"/><line x1="140" y1="236" x2="254.0862664711204" y2="163.74536456829043" stroke="#8b3a1f" stroke-width="1.4" stroke-linecap="round"/>
<polygon points="260,160 256.22647479585777,167.12464087050733 251.94605814638302,160.36608826607352" fill="#8b3a1f"/>
<!-- Broker -->
<g transform="translate(260,100)">
  <rect width="170" height="130" rx="12" fill="#dde7ec" stroke="#2a2622" stroke-width="1.5"/>
  <text x="85" y="36" text-anchor="middle" font-size="12" font-weight="700" fill="#2a5a7a">BROKER</text>
  <text x="85" y="60" text-anchor="middle" font-size="10" fill="#2a2622">topics</text>
  <text x="85" y="80" text-anchor="middle" font-size="10" font-family="monospace" fill="#6b6660">sensors/temp/#</text>
  <text x="85" y="96" text-anchor="middle" font-size="10" font-family="monospace" fill="#6b6660">sensors/co2/#</text>
  <text x="85" y="112" text-anchor="middle" font-size="10" font-family="monospace" fill="#6b6660">sensors/smoke/#</text>
</g>
<!-- Arrows out -->
<line x1="430" y1="165" x2="553.7390096630006" y2="103.1304951684997" stroke="#8b3a1f" stroke-width="1.4" stroke-linecap="round"/>
<polygon points="560,100 555.5278640450005,106.70820393249937 551.9501552810008,99.55278640450004" fill="#8b3a1f"/><line x1="430" y1="165" x2="553.0051717775808" y2="169.73096814529157" stroke="#8b3a1f" stroke-width="1.4" stroke-linecap="round"/>
<polygon points="560,170 552.851439289176,173.72801284381683 553.1589042659856,165.7339234467663" fill="#8b3a1f"/><line x1="430" y1="165" x2="553.9367008976598" y2="236.50194282557297" stroke="#8b3a1f" stroke-width="1.4" stroke-linecap="round"/>
<polygon points="560,240 551.9378110837015,239.96668516976737 555.935590711618,233.03720048137856" fill="#8b3a1f"/>
<!-- Consumers -->
  <g transform="translate(560,80)">
    <rect width="130" height="50" rx="8" fill="#e2ebde" stroke="#2a2622" stroke-width="1.5"/>
    <text x="65" y="24" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">DB writer</text>
    <text x="65" y="40" text-anchor="middle" font-size="9.5" fill="#6b6660">persist all</text>
  </g>
  <g transform="translate(560,150)">
    <rect width="130" height="50" rx="8" fill="#f4ead9" stroke="#2a2622" stroke-width="1.5"/>
    <text x="65" y="24" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">Anomaly model</text>
    <text x="65" y="40" text-anchor="middle" font-size="9.5" fill="#6b6660">streaming</text>
  </g>
  <g transform="translate(560,220)">
    <rect width="130" height="50" rx="8" fill="#f0d9d1" stroke="#2a2622" stroke-width="1.5"/>
    <text x="65" y="24" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">Dashboard</text>
    <text x="65" y="40" text-anchor="middle" font-size="9.5" fill="#6b6660">live view</text>
  </g>
<text x="345" y="265" text-anchor="middle" font-size="10" font-style="italic" fill="#6b6660">one source · many independent consumers</text>
</svg>
</div><figcaption>A broker in the middle decouples sensor producers from the consumers downstream. New consumers can subscribe without anyone changing the sensors.</figcaption></figure>

The first step in any pipeline is **ingestion**: getting data from sensors into a durable store. Three patterns dominate, each fitting a different scale.

### Pattern A: Direct REST POST

Each sensor sends every reading directly to an HTTP endpoint that writes it into a database immediately.

```
   sensor  ───POST /readings───►  database
```

Easy to set up. Works fine at low volume. Breaks at scale: 100 sensors at 10 Hz means a thousand HTTP requests per second, which overwhelms most databases. There is also no buffer — if the database is slow or temporarily down, readings are lost.

Analogy: handing every letter directly to the postman, who must be standing on your doorstep every time you want to send one. Fine if you write one letter a week. Disastrous if you run a mail-order business.

### Pattern B: Message queue as buffer

Each sensor publishes to a message broker. A separate consumer process reads from the broker and writes into the database.

```
   sensor 1 ──┐                                     ┌──► database
   sensor 2 ──┼──►   broker (MQTT/Kafka)   ──►  consumer
   sensor 3 ──┘                                     └──► (other consumers)
```

The broker acts as a shock absorber. If the database is briefly slow, messages accumulate in the broker, and the consumer catches up later. The sensor doesn't care whether the database is up.

This is the recommended pattern for the course: MQTT (using Mosquitto as the broker) plus a Python consumer that writes into TimescaleDB or DuckDB.

Analogy: the postman is replaced by a mailbox in the village square. You drop your letter in any time, the postman collects them on a schedule, the receiving end opens its mailbox when ready. Three actors, each independent.

### Pattern C: Batch file upload

The sensor accumulates readings in memory and writes them to a file (typically a Parquet file) once a minute or once an hour. A separate process reads those files into a database.

```
   sensor accumulates 60 s in memory  ──►  writes one file per minute  ──►  batch loader
```

Highest throughput, lowest network overhead. Bad for real-time decisions because of the latency, excellent for historical training data.

Analogy: instead of sending one letter at a time, you wait until end-of-day and send a single envelope containing all your letters from that day. Cheap and efficient when nothing is urgent. Hopeless when something is.

Most real systems use **two patterns at once**: the broker for the real-time path, batch files for the historical path. The same reading is written into both places, and each place serves a different audience.

---

## Part 4 — Stream Processing: Working on Data as It Flows

<figure class="diagram"><div class="dgm-frame">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 720 300" role="img" aria-label="Stream processing — rolling window" class="dgm">
<text x="360" y="30" text-anchor="middle" font-size="14" font-weight="600" fill="#2a2622">Stream processing acts on each reading as it arrives</text>
<!-- Timeline -->
<line x1="40" y1="150" x2="680" y2="150" stroke="#6b6660" stroke-width="1.5"/>
<polygon points="680,150 672,146 672,154" fill="#6b6660"/>
<text x="680" y="172" text-anchor="end" font-size="10" fill="#6b6660">time</text>
<!-- Dots on timeline -->
<circle cx="60" cy="137.69106782026958" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="60" y1="141.69106782026958" x2="60" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/><circle cx="86" cy="127.0343622399297" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="86" y1="131.0343622399297" x2="86" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/><circle cx="112" cy="116.5242500668093" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="112" y1="120.5242500668093" x2="112" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/><circle cx="138" cy="110.62357793840603" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="138" y1="114.62357793840603" x2="138" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/><circle cx="164" cy="111.58371493426253" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="164" y1="115.58371493426253" x2="164" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/><circle cx="190" cy="117.19975237774231" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="190" y1="121.19975237774231" x2="190" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/><circle cx="216" cy="117.56667198738626" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="216" y1="121.56667198738626" x2="216" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/><circle cx="242" cy="131.87911622016276" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="242" y1="135.87911622016276" x2="242" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/><circle cx="268" cy="134.63456497475886" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="268" y1="138.63456497475886" x2="268" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/><circle cx="294" cy="126.50066350388242" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="294" y1="130.50066350388244" x2="294" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/><circle cx="320" cy="125.77916331486284" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="320" y1="129.77916331486284" x2="320" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/><circle cx="346" cy="111.55590499704161" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="346" y1="115.55590499704161" x2="346" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/><circle cx="372" cy="118.93975903179386" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="372" y1="122.93975903179386" x2="372" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/><circle cx="398" cy="108.64861095322684" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="398" y1="112.64861095322684" x2="398" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/><circle cx="424" cy="116.10491830722312" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="424" y1="120.10491830722312" x2="424" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/><circle cx="450" cy="117.81720669365123" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="450" y1="121.81720669365123" x2="450" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/><circle cx="476" cy="111.82886435517696" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="476" y1="115.82886435517696" x2="476" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/><circle cx="502" cy="119.28605489134205" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="502" y1="123.28605489134205" x2="502" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/><circle cx="528" cy="126.7034703795506" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="528" y1="130.7034703795506" x2="528" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/><circle cx="554" cy="111.6536993070161" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="554" y1="115.6536993070161" x2="554" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/><circle cx="580" cy="135.71824416129854" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="580" y1="139.71824416129854" x2="580" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/><circle cx="606" cy="131.96205067462418" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="606" y1="135.96205067462418" x2="606" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/><circle cx="632" cy="125.7825230769656" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="632" y1="129.78252307696562" x2="632" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/><circle cx="658" cy="112.21007159271315" r="4" fill="#8b3a1f" opacity="0.85"/><line x1="658" y1="116.21007159271315" x2="658" y2="150" stroke="#8b3a1f" stroke-width="0.8" opacity="0.4"/>
<!-- Rolling window -->
<rect x="380" y="80" width="200" height="100" rx="10" fill="#f3ece6" stroke="#8b3a1f" stroke-width="2" opacity="0.8"/>
<text x="480" y="98" text-anchor="middle" font-size="11" font-weight="600" fill="#8b3a1f">rolling window</text>
<text x="480" y="114" text-anchor="middle" font-size="10" fill="#8b3a1f">last 60 seconds</text>
<!-- Window movement arrow -->
<line x1="580" y1="130" x2="633" y2="130" stroke="#8b3a1f" stroke-width="1.5" stroke-linecap="round"/>
<polygon points="640,130 633,134 633,126" fill="#8b3a1f"/>
<!-- Operator -->
<g transform="translate(280,210)">
  <rect width="200" height="50" rx="10" fill="#e2ebde" stroke="#2a2622" stroke-width="1.5"/>
  <text x="100" y="24" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">Operator</text>
  <text x="100" y="40" text-anchor="middle" font-size="10" fill="#6b6660">moving average · spike detection</text>
</g>
<line x1="480" y1="182" x2="386.7407481379993" y2="208.1125905213602" stroke="#8b3a1f" stroke-width="1.4" stroke-linecap="round"/>
<polygon points="380,210 385.6622284359194,204.2607344425035 387.81926784007914,211.96444660021692" fill="#8b3a1f"/>
<text x="360" y="290" text-anchor="middle" font-size="10" font-style="italic" fill="#6b6660">latency: milliseconds — answers appear before the data finishes arriving</text>
</svg>
</div><figcaption>Stream processing keeps a sliding window of recent readings in memory and produces answers continuously, with millisecond latency.</figcaption></figure>

Once data is flowing in, the next question is how to process it. Two main modes exist: streaming (processing each reading as it arrives) and batch (processing a large pile of stored data at once). This part covers streaming; the next covers batch.

A **stream** is an unbounded, time-ordered sequence of events. New events keep arriving forever. Stream processors apply operations to streams and produce either new streams or side effects (database writes, alerts, dashboard updates).

A useful image: a river flowing past a water-quality station. The station measures every drop as it goes past — there is no way to stop the river and inspect it all at once.

### Windowed aggregation

A single sensor reading is rarely useful by itself. It is noisy, missing context, and easily misleading. The fix is to compute statistics over a **window** of recent readings.

Three common window types:

**Tumbling windows** are non-overlapping, fixed-size buckets. Each reading falls into exactly one window.

```
   |----window 1----|----window 2----|----window 3----|
   |   00:00–00:05  |   00:05–00:10  |   00:10–00:15  |
   |                |                |                |
   readings...      readings...      readings...
```

Tumbling windows are useful for periodic reports: a 5-minute average that resets every 5 minutes.

**Sliding windows** are overlapping. Each new reading falls into multiple windows.

```
   |---window covers last 5 min---|
                |---window covers last 5 min---|
                          |---window covers last 5 min---|
   ...now                  ...now+1                ...now+2
```

Sliding windows are useful for continuous tracking: "the rolling 5-minute average, updated every second."

**Session windows** are dynamic. They grow as long as events keep arriving and close when there's a gap.

```
   |---session---|       gap        |---session---|
   read read read                   read read
```

Session windows are useful for grouping bursts of activity: "all access events for one person in one trip through the building."

A natural image: imagine watching a movie versus watching the news.
- Tumbling windows are like episodes — each one starts at a fixed time and runs a fixed length.
- Sliding windows are like a live ticker that always shows "the last 5 minutes."
- Session windows are like a phone call — they start when you pick up, end when you hang up, and have no fixed length.

### Complex event processing

Some patterns need to span multiple events over time. A fire detection rule might say: smoke above 0.7 for ten consecutive readings, combined with temperature rising at more than 2°C per minute, combined with a door opening in the same zone recently. No single reading triggers the rule; only the pattern across many readings does.

This is **complex event processing**, abbreviated CEP. Apache Flink has a CEP library that lets you describe these temporal patterns declaratively. For smaller systems, a stateful Python process maintaining a sliding window in memory is sufficient.

Analogy: a single yawn doesn't mean someone is bored. A yawn, followed by checking their phone, followed by looking at the exit, repeated three times in five minutes — that's a pattern. CEP is the skill of recognising the pattern, not the individual symptoms.

### Tools for stream processing

The right tool depends on scale.

**Apache Flink** is a distributed stream processing engine. Production-grade and feature-rich, but heavy to operate. Right when handling millions of events per second.

**Kafka Streams** is a stream processing library built directly on top of Apache Kafka. If Kafka is already in the architecture, Kafka Streams is the natural choice.

**Redis Streams** is a lightweight stream storage feature inside Redis. Built-in consumer groups, in-memory speed, easy to run.

**Plain Python with asyncio** is sufficient for the course. A simple consumer that reads from MQTT, computes windowed aggregations in memory, and writes to a database is entirely fine for one building.

---

## Part 5 — Batch Processing: Working on Data After It Settles

<figure class="diagram"><div class="dgm-frame">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 720 300" role="img" aria-label="Batch processing — accumulate, then process" class="dgm">
<text x="360" y="30" text-anchor="middle" font-size="14" font-weight="600" fill="#2a2622">Batch processing waits for data to settle</text>
<!-- Accumulating pile -->
<g transform="translate(60,90)">
  <text x="100" y="0" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">readings collect over 24 h</text>
  <rect x="0" y="30" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.58354050955375"/><rect x="22" y="30" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.7709718128537411"/><rect x="44" y="30" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.7215726175567279"/><rect x="66" y="30" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.7961893851263965"/><rect x="88" y="30" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.6096143530530314"/><rect x="110" y="30" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.7267102860373983"/><rect x="132" y="30" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.40352455880955557"/><rect x="154" y="30" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.7676489372213462"/><rect x="0" y="52" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.7235847980818126"/><rect x="22" y="52" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.6931860640129712"/><rect x="44" y="52" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.5171255960491383"/><rect x="66" y="52" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.5485025180034772"/><rect x="88" y="52" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.5057863351970471"/><rect x="110" y="52" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.6737490166833934"/><rect x="132" y="52" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.709200024276273"/><rect x="154" y="52" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.5435142031878195"/><rect x="0" y="74" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.8587639120211747"/><rect x="22" y="74" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.5152973979428195"/><rect x="44" y="74" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.8459559591059096"/><rect x="66" y="74" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.5380312845864945"/><rect x="88" y="74" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.5762637514330995"/><rect x="110" y="74" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.7724183237299822"/><rect x="132" y="74" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.5858081241263571"/><rect x="154" y="74" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.7054831018822191"/><rect x="0" y="96" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.8390222849035274"/><rect x="22" y="96" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.8456289260315373"/><rect x="44" y="96" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.8688018586332698"/><rect x="66" y="96" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.8184163412112202"/><rect x="88" y="96" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.6944401091490404"/><rect x="110" y="96" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.45881973307334767"/><rect x="132" y="96" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.7008670308171119"/><rect x="154" y="96" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.8139018300883945"/><rect x="0" y="118" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.5668999559238751"/><rect x="22" y="118" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.525986113932835"/><rect x="44" y="118" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.8523212324054299"/><rect x="66" y="118" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.7713871294324837"/><rect x="88" y="118" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.46364006968798993"/><rect x="110" y="118" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.7159823563562582"/><rect x="132" y="118" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.6987346145678066"/><rect x="154" y="118" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.6035633653706605"/><rect x="0" y="140" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.5607849888860854"/><rect x="22" y="140" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.5762522720347163"/><rect x="44" y="140" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.7823558529045014"/><rect x="66" y="140" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.5957910090488178"/><rect x="88" y="140" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.6263058693155441"/><rect x="110" y="140" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.7103859762084832"/><rect x="132" y="140" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.4389996080473376"/><rect x="154" y="140" width="18" height="18" rx="2" fill="#8b3a1f" opacity="0.6505474567343614"/>
</g>
<line x1="280" y1="150" x2="353" y2="150" stroke="#8b3a1f" stroke-width="2.2" stroke-linecap="round"/>
<polygon points="360,150 353,154 353,146" fill="#8b3a1f"/>
<text x="320" y="138" text-anchor="middle" font-size="10" font-weight="600" fill="#8b3a1f">2 a.m.</text>
<text x="320" y="170" text-anchor="middle" font-size="9.5" fill="#6b6660">scheduled run</text>
<!-- batch job box -->
<g transform="translate(380,110)">
  <rect width="160" height="80" rx="10" fill="#e2ebde" stroke="#2a2622" stroke-width="1.5"/>
  <text x="80" y="34" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">Batch job</text>
  <text x="80" y="52" text-anchor="middle" font-size="10" fill="#6b6660">aggregate · ML training</text>
  <text x="80" y="68" text-anchor="middle" font-size="10" fill="#6b6660">~30 minutes</text>
</g>
<line x1="540" y1="150" x2="603" y2="150" stroke="#8b3a1f" stroke-width="2.2" stroke-linecap="round"/>
<polygon points="610,150 603,154 603,146" fill="#8b3a1f"/>
<g transform="translate(610,130)">
  <rect width="80" height="40" rx="6" fill="#f4ead9" stroke="#2a2622" stroke-width="1.5"/>
  <text x="40" y="20" text-anchor="middle" font-size="11" font-weight="600" fill="#2a2622">Reports</text>
  <text x="40" y="34" text-anchor="middle" font-size="9.5" fill="#6b6660">models</text>
</g>
<text x="360" y="290" text-anchor="middle" font-size="10" font-style="italic" fill="#6b6660">latency: minutes to hours — for the answers that don't have to be live</text>
</svg>
</div><figcaption>Batch processing waits for data to settle, then processes large windows at once. Cheaper, simpler, but the freshest answer is from yesterday.</figcaption></figure>

Batch processing operates on large volumes of stored data all at once. A daily job that prepares training data, a weekly report that summarises energy use, a monthly anomaly analysis. Batch jobs run on a schedule (nightly, weekly) or on demand.

A useful image: streaming is fishing with a line — one fish at a time, you react as each one bites. Batch is fishing with a net — you collect a lot at once and process them together.

For building control, the critical batch job is **training data preparation**: pulling ninety days of sensor readings out of the data lake, computing the features the ML model needs (rolling averages, occupancy patterns, time-of-day encodings), and producing a clean Parquet or CSV file for training. This job may take minutes to hours, which is fine because real-time isn't required.

### Tools for batch processing

**DuckDB** is the recommended tool for the course. It is an in-process SQL engine that queries Parquet files directly without a server. Fast, zero-configuration, full SQL. It treats a folder of Parquet files like a database table.

```sql
-- DuckDB reading directly from a folder of Parquet files
SELECT room, AVG(value) AS avg_temp
FROM read_parquet('data/bronze/temperature/*.parquet')
WHERE ts > now() - INTERVAL '24 hours'
GROUP BY room;
```

**pandas** is a Python DataFrame library. Flexible, interactive, great for prototyping. Limited to a single machine and single thread, so it doesn't scale to industrial volumes, but for a building it is more than enough.

**Apache Spark** is a distributed batch processing engine. The right tool when one machine isn't enough — multi-building analytics at industrial scale.

---

## Part 6 — The Data Lake (with the Medallion Architecture)

<figure class="diagram"><div class="dgm-frame">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 720 320" role="img" aria-label="Medallion architecture — bronze, silver, gold" class="dgm">
<text x="360" y="30" text-anchor="middle" font-size="14" font-weight="600" fill="#2a2622">Data refines through three tiers</text>
  <g transform="translate(40,70)">
    <rect width="200" height="200" rx="14" fill="#f0d9b8" stroke="#2a2622" stroke-width="1.5"/>
    <circle cx="100" cy="60" r="34" fill="#d9a370" stroke="#2a2622" stroke-width="1.5"/>
    <text x="100" y="66" text-anchor="middle" font-size="13" font-weight="700" fill="#fdfbf7">B</text>
    <text x="100" y="118" text-anchor="middle" font-size="14" font-weight="700" fill="#2a2622">Bronze</text>
    <text x="100" y="136" text-anchor="middle" font-size="11" font-style="italic" fill="#6b6660">raw · unfiltered</text>
    <line x1="30" y1="152" x2="170" y2="152" stroke="#2a2622" stroke-width="0.6" opacity="0.4"/>
    <text x="100" y="172" text-anchor="middle" font-size="10" fill="#2a2622">every reading, exactly as it arrived</text>
  </g>
  <line x1="240" y1="170" x2="263" y2="170" stroke="#8b3a1f" stroke-width="1.8" stroke-linecap="round"/>
<polygon points="270,170 263,174 263,166" fill="#8b3a1f"/><text x="255" y="160" text-anchor="middle" font-size="9" fill="#8b3a1f">refine</text>
  <g transform="translate(260,70)">
    <rect width="200" height="200" rx="14" fill="#e6e6e6" stroke="#2a2622" stroke-width="1.5"/>
    <circle cx="100" cy="60" r="34" fill="#a8a8a8" stroke="#2a2622" stroke-width="1.5"/>
    <text x="100" y="66" text-anchor="middle" font-size="13" font-weight="700" fill="#fdfbf7">S</text>
    <text x="100" y="118" text-anchor="middle" font-size="14" font-weight="700" fill="#2a2622">Silver</text>
    <text x="100" y="136" text-anchor="middle" font-size="11" font-style="italic" fill="#6b6660">clean · validated</text>
    <line x1="30" y1="152" x2="170" y2="152" stroke="#2a2622" stroke-width="0.6" opacity="0.4"/>
    <text x="100" y="172" text-anchor="middle" font-size="10" fill="#2a2622">deduplicated, typed, joined</text>
  </g>
  <line x1="460" y1="170" x2="483" y2="170" stroke="#8b3a1f" stroke-width="1.8" stroke-linecap="round"/>
<polygon points="490,170 483,174 483,166" fill="#8b3a1f"/><text x="475" y="160" text-anchor="middle" font-size="9" fill="#8b3a1f">refine</text>
  <g transform="translate(480,70)">
    <rect width="200" height="200" rx="14" fill="#f4e5b8" stroke="#2a2622" stroke-width="1.5"/>
    <circle cx="100" cy="60" r="34" fill="#d9b34a" stroke="#2a2622" stroke-width="1.5"/>
    <text x="100" y="66" text-anchor="middle" font-size="13" font-weight="700" fill="#fdfbf7">G</text>
    <text x="100" y="118" text-anchor="middle" font-size="14" font-weight="700" fill="#2a2622">Gold</text>
    <text x="100" y="136" text-anchor="middle" font-size="11" font-style="italic" fill="#6b6660">curated · business-ready</text>
    <line x1="30" y1="152" x2="170" y2="152" stroke="#2a2622" stroke-width="0.6" opacity="0.4"/>
    <text x="100" y="172" text-anchor="middle" font-size="10" fill="#2a2622">aggregates · features · KPIs</text>
  </g>
</svg>
</div><figcaption>Bronze, silver, gold. Data gets more refined and more trustworthy at each tier — and the rules that produce each tier are versioned in code.</figcaption></figure>

So far the discussion has been about *moving* data and *processing* data. The next question is where to *store* it for the long term.

### The data lake philosophy

A **data lake** is a storage system that keeps raw data in its native format, exactly as it arrived, until somebody needs it. The philosophy is "store first, structure later." Unlike a data warehouse, which requires the schema to be decided up-front, a data lake lets the schema be decided at query time.

Why does this matter? Because in a real CPS, you don't know in week one that the 5-minute variance of CO2 readings will turn out to be a useful occupancy feature. By the time you discover it, you'd be furious if you had discarded that level of detail. The data lake keeps everything, in case you need it later.

The three core principles:

1. **Store raw data.** Never transform or discard the original sensor reading. Always keep the immutable raw record.
2. **Transform on read.** Apply cleaning, feature engineering, and aggregation at query time, not at write time. This way, if you find a bug in your transformation logic, you can fix it without re-collecting the data.
3. **Schema on read.** Define the data's shape when querying, not when storing. This accommodates evolving schemas without painful migrations.

Analogy: a data lake is like a pantry that stores raw ingredients — flour, eggs, vegetables — exactly as they arrived from the grocery store. A data warehouse is like a freezer full of pre-cooked meals: faster to serve, but if you decide tomorrow that you want to make a salad instead of lasagna, you're out of luck.

### The medallion architecture

Storing everything raw is helpful, but querying everything raw every time is expensive. The compromise is the **medallion architecture**, popularised by Databricks: organise the data lake into layers, each transformed a bit more than the previous one.

```
   BRONZE LAYER  ──►  SILVER LAYER  ──►  GOLD LAYER  ──►  MODEL ZONE
   ───────────       ────────────       ─────────       ───────────
   raw sensor        cleaned,            ML-ready        train/val/test
   readings,         deduplicated,       features:       datasets,
   exactly as        unified schema,     rolling avgs,   labelled,
   received,         missing values      time encodings, split
   IMMUTABLE         handled             cross-sensor
                                         relationships
```

**Bronze** is the raw zone. Each reading is stored exactly as it came off the wire, timestamped on arrival, and never modified. If a bug is found in a downstream pipeline, reprocessing from bronze is always possible.

**Silver** is the cleaned zone. Validated data (out-of-range values flagged), deduplicated (retransmissions removed), unified schema across sensors, missing values handled or marked. Silver is what most queries hit.

**Gold** is the features zone. Business-ready features for ML: rolling statistics, derived signals, time-of-day encodings, cross-sensor relationships. This is the input to model training.

**Model zone** holds ready-to-train datasets with labels and train/validation/test splits.

Analogy: bronze is your shopping receipts in a shoebox. Silver is the same receipts entered into a spreadsheet with consistent columns. Gold is a monthly summary showing categories, totals, and trends. Each is more useful than the last, but each was derived from the original receipts. If you discover an error in the gold report, you can always go back to the shoebox and reprocess.

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

The key insight is **columnar storage**. CSV stores data row by row:

```
   row 1: ts=…, sensor_id=…, room=…, value=…
   row 2: ts=…, sensor_id=…, room=…, value=…
   row 3: ts=…, sensor_id=…, room=…, value=…
```

Parquet stores it column by column:

```
   column ts:        [ts1, ts2, ts3, ts4, ts5, ts6, ts7, ts8, ...]
   column sensor_id: [s1,  s1,  s1,  s2,  s2,  s2,  s3,  s3,  ...]
   column room:      [r1,  r1,  r1,  r1,  r1,  r1,  r2,  r2,  ...]
   column value:     [v1,  v2,  v3,  v4,  v5,  v6,  v7,  v8,  ...]
```

This matters for two reasons. First, queries that need only the `value` column don't have to read the rest. Second, columnar data compresses much better than row-based data, because consecutive values are similar (e.g., the `room` column has just a few distinct strings repeated millions of times).

Parquet files are typically 4 to 10 times smaller than equivalent CSV files, and queries that touch only a few columns run an order of magnitude faster.

Analogy: a CSV is a stack of paper forms. Each form has multiple fields, and if you want to compute the average of one field across a thousand forms, you have to flip through every form. Parquet is a spreadsheet with one column per variable; computing the average of a column is one operation on one column.

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

**DuckDB** queries this directly using the `httpfs` extension. No ETL server, no cluster, no managed service.

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

<figure class="diagram"><div class="dgm-frame">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 720 280" role="img" aria-label="Time-series database" class="dgm">
<text x="360" y="30" text-anchor="middle" font-size="14" font-weight="600" fill="#2a2622">Time-series databases are tuned for one shape of data</text>
<!-- regular DB -->
<g transform="translate(60,80)">
  <rect width="200" height="160" rx="10" fill="#fdfbf7" stroke="#6b6660" stroke-width="1.5"/>
  <text x="100" y="26" text-anchor="middle" font-size="11" font-weight="700" fill="#6b6660">RELATIONAL DB</text>
  <g>
    <ellipse cx="100" cy="52" rx="30" ry="6" fill="#dde7ec" stroke="#2a2622" stroke-width="1.4"/>
    <path d="M 70 52 L 70 108 Q 100 116 130 108 L 130 52" fill="#dde7ec" stroke="#2a2622" stroke-width="1.4"/>
    <ellipse cx="100" cy="52" rx="30" ry="6" fill="none" stroke="#2a2622" stroke-width="1.4"/>
    <ellipse cx="100" cy="62" rx="28" ry="5" fill="none" stroke="#2a2622" stroke-width="0.8" opacity="0.5"/>
    <ellipse cx="100" cy="74" rx="27" ry="5" fill="none" stroke="#2a2622" stroke-width="0.8" opacity="0.5"/>
  </g>
  <text x="100" y="142" text-anchor="middle" font-size="10" fill="#2a2622">rows · keys · joins</text>
  <text x="100" y="156" text-anchor="middle" font-size="9.5" fill="#6b6660" font-style="italic">general-purpose</text>
</g>
<line x1="280" y1="160" x2="333" y2="160" stroke="#6b6660" stroke-width="1.5" stroke-linecap="round"/>
<polygon points="340,160 333,164 333,156" fill="#6b6660"/>
<text x="310" y="148" text-anchor="middle" font-size="10" fill="#6b6660">vs</text>
<!-- time-series DB -->
<g transform="translate(360,80)">
  <rect width="290" height="160" rx="10" fill="#e2ebde" stroke="#2a2622" stroke-width="1.5"/>
  <text x="145" y="26" text-anchor="middle" font-size="11" font-weight="700" fill="#3a5a3a">TIME-SERIES DB</text>
  <!-- chunks by time -->
  <rect x="20" y="50" width="36" height="60" rx="3" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/><rect x="60" y="50" width="36" height="60" rx="3" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/><rect x="100" y="50" width="36" height="60" rx="3" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/><rect x="140" y="50" width="36" height="60" rx="3" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/><rect x="180" y="50" width="36" height="60" rx="3" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/><rect x="220" y="50" width="36" height="60" rx="3" fill="#fdfbf7" stroke="#2a2622" stroke-width="1"/>
  <text x="38" y="120" text-anchor="middle" font-size="8" fill="#6b6660">09:00</text><text x="78" y="120" text-anchor="middle" font-size="8" fill="#6b6660">10:00</text><text x="118" y="120" text-anchor="middle" font-size="8" fill="#6b6660">11:00</text><text x="158" y="120" text-anchor="middle" font-size="8" fill="#6b6660">12:00</text><text x="198" y="120" text-anchor="middle" font-size="8" fill="#6b6660">13:00</text><text x="238" y="120" text-anchor="middle" font-size="8" fill="#6b6660">14:00</text>
  <circle cx="32.14493019872775" cy="60" r="2" fill="#8b3a1f"/><circle cx="31.812033719713373" cy="72" r="2" fill="#8b3a1f"/><circle cx="34.496815568083136" cy="84" r="2" fill="#8b3a1f"/><circle cx="36.30165808028192" cy="96" r="2" fill="#8b3a1f"/><circle cx="72.46711143593049" cy="60" r="2" fill="#8b3a1f"/><circle cx="76.16842794430146" cy="72" r="2" fill="#8b3a1f"/><circle cx="70.34742750452574" cy="84" r="2" fill="#8b3a1f"/><circle cx="79.83153278802558" cy="96" r="2" fill="#8b3a1f"/><circle cx="119.29248020556241" cy="60" r="2" fill="#8b3a1f"/><circle cx="113.31128127293475" cy="72" r="2" fill="#8b3a1f"/><circle cx="111.7177515844459" cy="84" r="2" fill="#8b3a1f"/><circle cx="111.57450501107621" cy="96" r="2" fill="#8b3a1f"/><circle cx="156.98920883510516" cy="60" r="2" fill="#8b3a1f"/><circle cx="159.32343746297823" cy="72" r="2" fill="#8b3a1f"/><circle cx="160.26684259566184" cy="84" r="2" fill="#8b3a1f"/><circle cx="158.37638349111023" cy="96" r="2" fill="#8b3a1f"/><circle cx="191.04786647113912" cy="60" r="2" fill="#8b3a1f"/><circle cx="195.75286568079974" cy="72" r="2" fill="#8b3a1f"/><circle cx="203.16823703305926" cy="84" r="2" fill="#8b3a1f"/><circle cx="192.84578089867776" cy="96" r="2" fill="#8b3a1f"/><circle cx="243.95115354373297" cy="60" r="2" fill="#8b3a1f"/><circle cx="232.4105862145665" cy="72" r="2" fill="#8b3a1f"/><circle cx="228.3429618317438" cy="84" r="2" fill="#8b3a1f"/><circle cx="241.31790943336227" cy="96" r="2" fill="#8b3a1f"/>
  <text x="145" y="144" text-anchor="middle" font-size="10" fill="#2a2622">chunks by time · compressed · auto-aged</text>
  <text x="145" y="158" text-anchor="middle" font-size="9.5" fill="#6b6660" font-style="italic">TimescaleDB · InfluxDB · QuestDB</text>
</g>
</svg>
</div><figcaption>Time-series databases store rows indexed by time, compressed per chunk, and discarded automatically after a retention period. Far faster than a general-purpose database for this shape of data.</figcaption></figure>

A general-purpose database like PostgreSQL or MySQL is built for a typical web-application workload: lots of lookups by primary key, joins between tables, transactional updates. Sensor data has a completely different shape.

### Why time-series data is different

Three properties of sensor data that don't fit a general database:

1. **Append-heavy.** Every write is a new row. Existing rows are never updated. A general database's update machinery is unused overhead.
2. **Time-range queries dominate.** Almost every question is "what happened between time A and time B?" Joining tables by primary key is rare.
3. **Downsampling.** Two-year-old 5-second data is rarely needed at full resolution. Hourly averages would do, freeing 99% of the storage.

A **time-series database**, abbreviated TSDB, is built for exactly this shape. Data is stored in time-ordered chunks. New data appends to the latest chunk. Time-range predicates immediately prune entire chunks outside the range. Retention policies automatically downsample or delete old data.

Analogy: a general-purpose database is like a filing cabinet where every paper is filed by topic. Finding "all papers about Project X" is fast; finding "all papers from last March" requires walking through every folder. A time-series database is like a diary, where every page is dated and finding "what happened in March" is just opening to March.

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

**ClickHouse** is a column-oriented analytical database, extremely fast for read-heavy analytical queries. Used in production at Cloudflare and Uber. Appropriate when analytics push beyond what TimescaleDB can comfortably handle.

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

A useful image: feature engineering is what a chef does to crude ingredients before serving. The raw potato is not edible. Wash, peel, slice, and bake it, and now you have something useful.

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

Analogy: imagine writing the hours on a clock face. The model that uses the *clock face position* (which is what sin/cos gives it) understands that 11 o'clock and 1 o'clock are nearby. The model that uses *just the number* sees 11 and 1 as nine hours apart.

### Cross-sensor features

Relationships between sensors are often more informative than any single sensor. A few examples:

- Temperature difference between adjacent rooms — detects a door left open or an HVAC imbalance.
- CO2 combined with ventilation state — estimates occupancy without an occupancy sensor.
- Smoke level combined with temperature gradient — distinguishes cooking (high smoke, modest temperature rise) from a real fire (high smoke, fast temperature rise).

### Load

Writing the computed features to a feature store — a database or Parquet file ready for ML training — or feeding them directly into the model.

### ETL tools

**dbt (data build tool)** is a popular open-source tool that lets transformations be expressed as SQL queries. Dbt runs queries in the correct order, tests their outputs, and generates documentation. Free, well-supported, with a free 4-hour fundamentals course.

**pandas** is the standard Python DataFrame library. Flexible and interactive, ideal for exploratory work. Single-threaded and in-memory, so it doesn't scale to industrial volumes; production pipelines are usually rewritten in SQL (using dbt and DuckDB) once the transformations are stable.

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

**Clock drift.** Each sensor has its own clock, and clocks drift. A reading from sensor A timestamped 14:32:00 and one from sensor B timestamped 14:31:58 may have happened in the opposite order from what the timestamps suggest. Mitigations: run NTP (Network Time Protocol) on every device, store both the device timestamp and the time when the server received the message, and prefer the server timestamp for ordering.

Analogy: imagine a courtroom where every witness uses a different clock. Their statements about timing can't be compared without first synchronising the clocks.

**Schema evolution.** A new sensor type is added with an extra field. The existing pipeline doesn't know what to do with that field. Schema-on-read (typical for data lakes) handles this gracefully — old code ignores new fields. Schema-on-write (typical for relational databases) requires a migration to add the column.

### Monitoring the pipeline

A data pipeline that silently fails is worse than one that fails loudly. Production-quality pipelines instrument themselves with metrics, and a monitoring system alerts when something looks wrong.

Four metrics worth tracking:

- **Data freshness.** Age of the most recent reading for each sensor. Alert if it exceeds three times the expected interval.
- **Value range.** A temperature below -50°C or above 100°C inside a Swedish office building is impossible. Alert when readings violate physical bounds.
- **Volume.** Number of readings per minute. A sudden drop signals offline sensors or a broken consumer.
- **Error rate.** Number of readings rejected by validation. A spike suggests a schema change or a faulty sensor.

The standard stack for this is **Prometheus** (a metrics database that scrapes endpoints periodically) plus **Grafana** (a dashboard tool that visualises the metrics and triggers alerts). Both are free, both run in Docker, both are industry standard.

---

## Part 12 — A Worked Example: Smoke Detection Data Pipeline

<figure class="diagram"><div class="dgm-frame">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 720 360" role="img" aria-label="Smoke detection data pipeline" class="dgm">
<text x="360" y="30" text-anchor="middle" font-size="14" font-weight="600" fill="#2a2622">End-to-end pipeline for smoke detection</text>
  <g transform="translate(30,90)">
    <rect width="100" height="110" rx="10" fill="#f3ece6" stroke="#2a2622" stroke-width="1.5"/>
    <g transform="translate(50,42)"><g transform="translate(0,0)">
    <circle cx="0" cy="0" r="13" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.4"/>
    <circle cx="0" cy="0" r="3.9000000000000004" fill="#8b3a1f"/>
    <path d="M -10.4 -13 Q -15.600000000000001 -18.2 -10.4 -23.400000000000002" fill="none" stroke="#6b6660" stroke-width="1.2"/>
    <path d="M 10.4 -13 Q 15.600000000000001 -18.2 10.4 -23.400000000000002" fill="none" stroke="#6b6660" stroke-width="1.2"/>
  </g></g>
    <text x="50" y="88" text-anchor="middle" font-size="10.5" font-weight="600" fill="#2a2622">Smoke sensor</text>
    <text x="50" y="102" text-anchor="middle" font-size="9" fill="#6b6660">10 Hz reading</text>
  </g>
  <g transform="translate(170,90)">
    <rect width="100" height="110" rx="10" fill="#dde7ec" stroke="#2a2622" stroke-width="1.5"/>
    <rect x="22" y="36" width="56" height="32" rx="6" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.3"/><text x="50" y="58" text-anchor="middle" font-size="10" font-weight="700" fill="#2a2622">MQTT</text>
    <text x="50" y="88" text-anchor="middle" font-size="10.5" font-weight="600" fill="#2a2622">Broker</text>
    <text x="50" y="102" text-anchor="middle" font-size="9" fill="#6b6660">MQTT topic</text>
  </g>
  <g transform="translate(310,90)">
    <rect width="100" height="110" rx="10" fill="#e2ebde" stroke="#2a2622" stroke-width="1.5"/>
    <polyline points="20,46 32,38 44,52 56,32 68,46 80,30" fill="none" stroke="#8b3a1f" stroke-width="1.8"/>
    <text x="50" y="88" text-anchor="middle" font-size="10.5" font-weight="600" fill="#2a2622">Stream processor</text>
    <text x="50" y="102" text-anchor="middle" font-size="9" fill="#6b6660">60s window</text>
  </g>
  <g transform="translate(450,90)">
    <rect width="100" height="110" rx="10" fill="#f4ead9" stroke="#2a2622" stroke-width="1.5"/>
    <g transform="translate(50,48)"><polyline points="-26,-4 -14,-8 -2,6 10,-16 22,-2 30,4" fill="none" stroke="#8b3a1f" stroke-width="2"/><circle cx="10" cy="-16" r="3" fill="#8b3a1f"/></g>
    <text x="50" y="88" text-anchor="middle" font-size="10.5" font-weight="600" fill="#2a2622">Anomaly model</text>
    <text x="50" y="102" text-anchor="middle" font-size="9" fill="#6b6660">scores spike</text>
  </g>
  <g transform="translate(590,90)">
    <rect width="100" height="110" rx="10" fill="#f0d9d1" stroke="#2a2622" stroke-width="1.5"/>
    <g transform="translate(50,46)"><g><circle cx="0" cy="0" r="18" fill="#8b3a1f" opacity="0.9"/>
    <text x="0" y="5" text-anchor="middle" font-size="12.6" font-weight="700" fill="#fdfbf7">AI</text></g></g>
    <text x="50" y="88" text-anchor="middle" font-size="10.5" font-weight="600" fill="#2a2622">Alert + log</text>
    <text x="50" y="102" text-anchor="middle" font-size="9" fill="#6b6660">safety agent</text>
  </g>
<line x1="130" y1="145" x2="163" y2="145" stroke="#8b3a1f" stroke-width="1.6" stroke-linecap="round"/>
<polygon points="170,145 163,149 163,141" fill="#8b3a1f"/><line x1="270" y1="145" x2="303" y2="145" stroke="#8b3a1f" stroke-width="1.6" stroke-linecap="round"/>
<polygon points="310,145 303,149 303,141" fill="#8b3a1f"/><line x1="410" y1="145" x2="443" y2="145" stroke="#8b3a1f" stroke-width="1.6" stroke-linecap="round"/>
<polygon points="450,145 443,149 443,141" fill="#8b3a1f"/><line x1="550" y1="145" x2="583" y2="145" stroke="#8b3a1f" stroke-width="1.6" stroke-linecap="round"/>
<polygon points="590,145 583,149 583,141" fill="#8b3a1f"/>
<!-- Bottom: storage tier -->
<g transform="translate(40,250)">
  <rect width="640" height="80" rx="10" fill="#fdfbf7" stroke="#2a2622" stroke-width="1.5" stroke-dasharray="6 4"/>
  <text x="20" y="22" font-size="11" font-weight="700" letter-spacing="1.5" fill="#6b6660">STORAGE TIER</text>
  <g transform="translate(110,40)"><g>
    <ellipse cx="0" cy="-14" rx="23" ry="6" fill="#dde7ec" stroke="#2a2622" stroke-width="1.4"/>
    <path d="M -23 -14 L -23 14 Q 0 22 23 14 L 23 -14" fill="#dde7ec" stroke="#2a2622" stroke-width="1.4"/>
    <ellipse cx="0" cy="-14" rx="23" ry="6" fill="none" stroke="#2a2622" stroke-width="1.4"/>
    <ellipse cx="0" cy="-4" rx="21" ry="5" fill="none" stroke="#2a2622" stroke-width="0.8" opacity="0.5"/>
    <ellipse cx="0" cy="8" rx="20" ry="5" fill="none" stroke="#2a2622" stroke-width="0.8" opacity="0.5"/>
  </g></g>
  <text x="120" y="70" text-anchor="middle" font-size="9.5" fill="#2a2622">time-series DB</text>
  <g transform="translate(250,42)"><rect width="50" height="38" rx="4" fill="#f3ece6" stroke="#2a2622" stroke-width="1.2"/><text x="25" y="24" text-anchor="middle" font-size="9" fill="#2a2622">parquet</text></g>
  <text x="275" y="92" text-anchor="middle" font-size="9.5" fill="#2a2622">data lake</text>
  <g transform="translate(430,42)"><rect width="50" height="38" rx="4" fill="#e2ebde" stroke="#2a2622" stroke-width="1.2"/><text x="25" y="20" text-anchor="middle" font-size="8" fill="#2a2622">model</text><text x="25" y="32" text-anchor="middle" font-size="8" fill="#2a2622">artifacts</text></g>
  <text x="455" y="92" text-anchor="middle" font-size="9.5" fill="#2a2622">model registry</text>
</g>
<!-- arrows from sensors layer to storage -->
<line x1="220" y1="200" x2="165.37754895718163" y2="245.51870920234865" stroke="#6b6660" stroke-width="1.2" stroke-linecap="round"/>
<polygon points="160,250 162.81681135852372,242.44582408395914 167.93828655583954,248.59159432073815" fill="#6b6660"/>
<line x1="360" y1="200" x2="281.0335395061044" y2="246.45085911405624" stroke="#6b6660" stroke-width="1.2" stroke-linecap="round"/>
<polygon points="275,250 279.0054589998508,243.00312225342518 283.06162001235793,249.8985959746873" fill="#6b6660"/>
<line x1="500" y1="200" x2="459.68275312135717" y2="244.79694097626984" stroke="#6b6660" stroke-width="1.2" stroke-linecap="round"/>
<polygon points="455,250 456.7095765363685,242.12108204978003 462.65592970634583,247.47279990275965" fill="#6b6660"/>
</svg>
</div><figcaption>Every component of the smoke-detection data pipeline: sensors at the top push readings down through ingestion, processing, modelling, and alerting — and everything is persisted to storage for replay and training.</figcaption></figure>

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
| **Complex event processing (CEP)** | Detecting temporal patterns across multiple events |
| **Data lake** | A storage system that keeps raw data in its native format until it is needed |
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
