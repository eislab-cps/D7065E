# Use Case Catalogue

Choose one use case or propose your own (must be approved). All use cases require a **data-driven approach** — statistical methods, ML models, or AI reasoning. Rule-based thresholds alone are not sufficient.

---

## Safety & Emergency

### Fire Detection and Autonomous Response

Deploy smoke, CO, and temperature sensors across floors. Train an anomaly detection model on normal sensor readings to distinguish real fires from false alarms (cooking, dust, sensor malfunction). When a fire is detected, the system autonomously activates sprinklers, unlocks fire doors, computes evacuation routes via the walkable navigation graph, tracks person locations, and highlights escape routes on the 3D map. The system predicts fire spread direction using temperature gradient analysis across adjacent rooms.

**What you build:**

| Component | Technology |
|-----------|-----------|
| Smoke sensor process (per room) | Container, REST client, simulated smoke readings with realistic noise patterns |
| CO sensor process (per room) | Container, REST client, correlated with smoke for multi-sensor fusion |
| Temperature sensor process | Container, generates thermal gradients that propagate between adjacent rooms |
| Sprinkler actuator process | Container, receives activation commands, reports state |
| Fire door actuator process | Container, locks/unlocks doors, reports position |
| Anomaly detection model | Autoencoder or isolation forest trained on "normal" multi-sensor time series |
| Fire spread predictor | Spatial model using room adjacency graph + temperature gradients |
| Evacuation route planner | Uses BuildSim walkable graph API with dynamic re-routing around affected zones |
| AI agent | LangChain/LangGraph agent that orchestrates detection → response → evacuation |
| Data pipeline | Sensor readings → time-series DB, feature extraction for ML training |
| Dashboard | BuildSim session API: room highlights (red=fire, orange=risk, green=safe), coverage zones for smoke spread, route display for evacuation paths |

### Gas Leak Detection and Containment

Gas leak and air quality sensors monitor for hazardous concentrations. A predictive model estimates leak propagation speed and direction based on ventilation patterns and room connectivity. HVAC is shut down or redirected to prevent gas spread. Doors are locked/unlocked to create containment zones and evacuation corridors. Coverage zones visualize the predicted contamination area in real-time.

**What you build:**

| Component | Technology |
|-----------|-----------|
| Gas leak sensor processes | Containers simulating gas concentration readings with diffusion model |
| Air quality sensor processes | CO2, VOC sensors that react to gas propagation |
| HVAC actuator processes | Fan speed, damper position control |
| Door lock actuator processes | Zone containment via door control |
| Gas dispersion model | Graph neural network or physics-informed ML using room connectivity + ventilation topology |
| AI agent | Predicts affected rooms, commands HVAC shutdown, creates containment zones |
| Data pipeline | Real-time stream processing for fast detection, batch for model training |
| Dashboard | Coverage zones showing predicted contamination, room highlights for containment status |

---

## Energy & Sustainability

### Smart HVAC Optimization

Occupancy counters, temperature sensors, and humidity sensors feed an ML model that predicts room-by-room occupancy and thermal load. The system autonomously adjusts heating, cooling, and ventilation to minimize energy consumption while maintaining comfort. It learns daily/weekly occupancy patterns, predicts occupancy 30-60 minutes ahead for pre-conditioning, and balances comfort constraints against energy cost.

**What you build:**

| Component | Technology |
|-----------|-----------|
| Temperature sensor processes (per room) | Containers with thermal models (heat gain from occupants, loss through walls) |
| Humidity sensor processes | Correlated with occupancy and HVAC state |
| Occupancy counter processes | Simulated entry/exit counts with daily/weekly patterns |
| HVAC actuator processes | Heating, cooling, fan speed control per zone |
| Occupancy forecasting model | LSTM, Prophet, or transformer for time-series prediction |
| Thermal model | Predicts temperature response to HVAC settings + occupancy |
| Optimization agent | Reinforcement learning or model predictive control for HVAC scheduling |
| Energy tracker | Computes energy consumption, compares against fixed-schedule baseline |
| Data pipeline | Sensor readings → data lake (Parquet), feature engineering (rolling averages, time features), model training pipeline |
| Dashboard | Energy consumption graphs (Grafana), room temperature heatmap (BuildSim highlights), occupancy prediction vs. actual |

### Intelligent Lighting Control

Light sensors, occupancy sensors, and window orientation data drive an ML model that optimizes lighting levels per room. Predicts natural light availability based on time, season, and weather. Learns user preferences for different room types. Dims or turns off lights in unoccupied rooms with prediction-based pre-activation. Estimates and reports energy savings.

**What you build:**

| Component | Technology |
|-----------|-----------|
| Light level sensor processes | Simulate lux readings based on time-of-day, window orientation, weather |
| Occupancy sensor processes | PIR-style presence detection with entry/exit patterns |
| Light fixture actuator processes | Dimming control (0-100%), on/off |
| Daylight prediction model | Regression model using time, season, weather, room orientation |
| Occupancy prediction model | Learns daily patterns, predicts when rooms will be occupied |
| Lighting control agent | RL agent that balances natural light, artificial light, occupancy prediction, and energy cost |
| Data pipeline | Light + occupancy readings → feature store, energy consumption tracking |
| Dashboard | Room brightness visualization, energy savings report, prediction accuracy |

---

## Security

### Anomaly-Based Intrusion Detection

Motion sensors, door locks, and card readers generate access events. An ML model learns normal access patterns (who goes where, when) and flags anomalies. Builds behavioral profiles with time-of-day patterns, common paths, and dwell times. Detects anomalies: access at unusual hours, unusual room sequences, tailgating patterns. AI classifies alerts by severity and suggests response.

**What you build:**

| Component | Technology |
|-----------|-----------|
| Card reader sensor processes | Simulate badge-in/badge-out events with realistic patterns (employees, visitors) |
| Motion sensor processes | Detect presence, report room occupancy changes |
| Door lock actuator processes | Lock/unlock zones in response to security agent |
| Access pattern model | Clustering (DBSCAN, GMM) on access event sequences to build "normal" profiles |
| Anomaly detector | Isolation forest or one-class SVM scoring each access event in real-time |
| Path analyzer | Sequence model (HMM, LSTM) for detecting unusual room-to-room transitions |
| Security agent | LLM-based agent that receives anomaly alerts, assesses severity, recommends action |
| Data pipeline | Event stream → Kafka/MQTT, access logs → data lake, real-time scoring |
| Dashboard | BuildSim highlights for threat zones, anomalous paths drawn on map, alert feed |

---

## Network & Infrastructure

### WiFi Coverage Optimization

WiFi access points with signal strength sensors deployed across the building. An ML model predicts coverage quality and identifies dead zones. Builds a signal propagation model from sensor measurements. Predicts coverage impact of adding/moving access points. Optimizes AP power levels to balance coverage vs. interference. Visualizes coverage as spherical zones on the 3D map.

**What you build:**

| Component | Technology |
|-----------|-----------|
| WiFi AP equipment (per AP) | Register as `wifi_access_point` in BuildSim, signal strength + connected user sensors |
| Client probe sensor processes | Simulate signal strength measurements at various room locations |
| AP power actuator processes | Adjust transmit power per AP |
| Signal propagation model | Gaussian process regression or neural network mapping position → signal strength |
| Coverage optimizer | Bayesian optimization or genetic algorithm for AP power tuning |
| Congestion predictor | Time-series model predicting connected-user counts from occupancy data |
| AI agent | Monitors coverage quality, adjusts AP power, recommends new AP placements |
| Data pipeline | Signal measurements → data lake, spatial feature engineering |
| Dashboard | BuildSim coverage zones (spheres colored by signal quality), dead zone highlights |

### Predictive Maintenance

Vibration, temperature, and power sensors on equipment (compressors, pumps, HVAC units) feed a predictive model. Learns normal operating signatures per equipment type. Detects degradation trends. Predicts remaining useful life and optimal maintenance timing. Prioritizes maintenance routes using the navigation graph.

**What you build:**

| Component | Technology |
|-----------|-----------|
| Vibration sensor processes | Simulate vibration readings with gradual degradation patterns |
| Equipment temperature sensors | Monitor operating temperature (distinct from room temperature) |
| Power consumption sensors | Track kW per equipment unit |
| Equipment health model | Autoencoder trained on normal operating data, anomaly = degradation |
| Remaining useful life predictor | Survival analysis or regression on degradation trend |
| Failure mode classifier | Random forest classifying failure type from sensor signature |
| Maintenance agent | Plans inspection routes via walkable graph, prioritizes by urgency |
| Data pipeline | High-frequency sensor data → downsampled to features → model training |
| Dashboard | Equipment health status on 3D map (green/yellow/red), maintenance schedule, RUL predictions |

---

## Comfort & Wellbeing

### Indoor Air Quality Management

CO2, humidity, VOC, and temperature sensors monitor air quality per room. An ML model predicts air quality degradation based on occupancy and ventilation state. Proactively increases ventilation before quality drops below thresholds. Balances air quality vs. energy cost. Reports compliance with Swedish workplace regulations (AFS 2020:1).

**What you build:**

| Component | Technology |
|-----------|-----------|
| CO2 sensor processes (per room) | Simulate CO2 buildup correlated with occupancy (each person ~40 ppm/hour in typical room) |
| VOC sensor processes | Volatile organic compound levels linked to activities and materials |
| Humidity sensor processes | Related to occupancy, HVAC state, weather |
| Ventilation actuator processes | Fan speed, damper position, fresh air intake control |
| CO2 trajectory model | Predicts CO2 concentration 30 min ahead given current occupancy + ventilation |
| Ventilation control agent | RL or MPC agent balancing air quality targets vs. energy cost |
| Compliance checker | Monitors against AFS 2020:1 thresholds, flags predicted violations |
| Data pipeline | Sensor readings → time-series DB, rolling averages, compliance reports |
| Dashboard | Room-by-room air quality heatmap (BuildSim highlights), compliance status, energy vs. quality trade-off graph |

---

## Building Regulation Compliance

### Automated Regulatory Compliance Monitoring

An AI agent continuously monitors building state against regulatory requirements and predicts violations before they occur. Covers fire code (door closing times, sprinkler coverage, evacuation routes), ventilation regulations (BBR), temperature regulations (AFS 2020:1), legionella prevention (hot water >60°C), and radon levels (Swedish limit 200 Bq/m³).

**What you build:**

| Component | Technology |
|-----------|-----------|
| Fire door timing sensors | Measure door close time, compare against code requirements |
| Water temperature sensors | Monitor hot water pipes for legionella compliance |
| Radon sensor processes | Simulate seasonal radon variations (higher in winter, varies by floor) |
| Room temperature sensors | Workplace temperature compliance during work hours |
| Ventilation flow sensors | Air changes per hour per room type |
| Violation prediction model | Time-series forecasting: "temperature will breach minimum in 45 min" |
| Severity classifier | Classifies violations by risk level and regulatory source |
| Compliance agent | LLM agent that interprets regulations, monitors sensor data, takes preventive action |
| Data pipeline | Continuous sensor data → compliance event log → violation history → trend analysis |
| Dashboard | Compliance dashboard with red/yellow/green per regulation, violation timeline, predictive alerts on BuildSim map |

---

## Environmental Monitoring

### Mold Prevention System

Humidity and temperature sensors in at-risk areas (basements, bathrooms, exterior walls) feed a model that predicts condensation risk. Calculates dew point, predicts mold growth risk using established models (VTT model), adjusts ventilation and heating to keep surfaces above dew point. Learns building-specific thermal bridge locations from sensor data.

**What you build:**

| Component | Technology |
|-----------|-----------|
| Humidity sensor processes | Per-room humidity with wall proximity variants |
| Temperature sensor processes | Air temperature + simulated surface temperature near thermal bridges |
| Ventilation actuator processes | Dehumidification, fresh air control |
| Heating actuator processes | Radiator control near at-risk walls |
| Dew point calculator | Physics-based: `Td = T - (100 - RH) / 5` (Magnus approximation) |
| Mold growth risk model | VTT model parameterized with ML-learned building-specific factors |
| Thermal bridge detector | Learns locations where surface temp drops below dew point from historical data |
| Prevention agent | Adjusts ventilation and heating to maintain surface temp > dew point + margin |
| Data pipeline | Humidity + temperature → dew point computation → risk scoring → action |
| Dashboard | Risk heatmap on BuildSim (coverage zones for moisture risk), thermal bridge locations, seasonal trend graphs |

---

## Health & Pandemic Preparedness

### Airborne Infection Risk Management

CO2 sensors serve as proxy for aerosol concentration. The system models infection transmission risk per room based on occupancy, ventilation rate, and duration. Implements Wells-Riley or similar epidemiological model. AI agent manages occupancy limits and ventilation. Simulates "what-if" scenarios. Visualizes risk zones.

**What you build:**

| Component | Technology |
|-----------|-----------|
| CO2 sensor processes | Proxy for air freshness and aerosol concentration |
| Occupancy counter processes | Track how many people and for how long |
| Ventilation actuator processes | Control fresh air rate |
| Wells-Riley infection model | Epidemiological model: P(infection) = f(quanta rate, ventilation, duration, occupants) |
| Risk calibration model | ML-calibrated parameters from CO2 + occupancy data |
| Ventilation/occupancy agent | RL agent: minimize infection risk while maximizing usable room capacity |
| What-if simulator | "If person in A2306 is infected, what's the risk in adjacent rooms over next 2 hours?" |
| Data pipeline | CO2 + occupancy → risk computation → room-level risk scores |
| Dashboard | Coverage zones showing infection risk levels, room highlights, what-if scenario results |

### Contact Tracing Simulation

Track simulated person movements through the building. When a person is flagged as sick, compute exposure risk for all other persons. Simulate realistic movement patterns using the walkable navigation graph.

**What you build:**

| Component | Technology |
|-----------|-----------|
| Person movement simulator | Generates realistic paths using walkable graph (office → meeting → lunch → office) |
| Occupancy tracker | Records who was in which room at what time |
| Exposure risk model | Graph-based risk propagation: shared room time × ventilation quality |
| Movement predictor | Markov chain or LSTM predicting where people will go next |
| Contact tracing agent | Given "person X is sick at time T", compute exposure graph for all contacts |
| Data pipeline | Movement events → temporal graph database → exposure queries |
| Dashboard | Person paths on 3D map, exposure risk highlights, contact network visualization |

---

## Digital Twin & Observability

### Sensor Drift and Anomaly Detection

An AI model learns expected correlations between sensors (heater on → temperature rises, door opens → CO2 drops temporarily). It flags sensors whose readings deviate from expected behavior. Distinguishes sensor failure from real events. Recommends sensor maintenance or recalibration.

**What you build:**

| Component | Technology |
|-----------|-----------|
| Diverse sensor processes | Temperature, humidity, CO2, door, HVAC — some deliberately injecting drift or faults |
| Multi-sensor correlation model | VAR (vector autoregression) or graph neural network learning normal sensor relationships |
| Drift detector | Monitors model residuals — gradual increase in prediction error = drift |
| Fault classifier | Distinguishes stuck sensor, noisy sensor, drifting sensor, real anomaly |
| Sensor health agent | Assesses sensor fleet health, recommends maintenance, adjusts confidence in sensor readings |
| Data pipeline | Multi-sensor time series → feature computation → model training → real-time scoring |
| Dashboard | Sensor health map on BuildSim (green=healthy, yellow=drifting, red=faulty), drift trend graphs |

---

## Multi-Agent Systems

### Competing Objectives with Multi-Agent Coordination

Deploy multiple AI agents with different objectives that must negotiate when they conflict. Energy agent minimizes power consumption. Comfort agent maintains temperature, air quality, and lighting. Safety agent ensures fire code compliance and evacuation readiness. Security agent monitors access patterns.

**What you build:**

| Component | Technology |
|-----------|-----------|
| Full sensor suite | Temperature, CO2, humidity, smoke, motion, card readers, light, power — all as containers |
| Full actuator suite | HVAC, lighting, door locks, sprinklers — all as containers |
| Energy agent | Forecasting model + optimization: minimize kWh while respecting constraints |
| Comfort agent | Occupancy prediction + comfort model: maintain temperature/CO2/lighting targets |
| Safety agent | Anomaly detection + rule engine: monitor fire code compliance, always-on |
| Security agent | Access pattern model + anomaly detection: monitor for intrusions |
| Coordination layer | Priority-based (safety > comfort > energy) or auction/negotiation protocol |
| Conflict resolution | What happens when energy wants HVAC off but comfort wants it on? Logged and explained. |
| Data pipeline | Shared data lake, each agent has its own feature pipeline |
| Dashboard | Multi-agent decision log, current agent states, conflict history, building state on 3D map |
