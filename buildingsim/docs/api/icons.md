# Icons API

SVG equipment icons embedded in the binary. Used by the browser to render equipment on the map.

## Get Icon

```
GET /api/icons/{name}.svg
```

Returns the SVG file with `Content-Type: image/svg+xml`.

```bash
# Download an icon
curl http://localhost:9090/api/icons/compressor.svg -o compressor.svg

# View in browser
open http://localhost:9090/api/icons/alien.svg
```

## Icon Gallery

### Access Control
| | |
|:---:|:---:|
| ![door_lock](../../data/equipment/icons/door_lock.svg) | ![card_reader](../../data/equipment/icons/card_reader.svg) |
| `door_lock` | `card_reader` |

### HVAC & Climate
| | | | |
|:---:|:---:|:---:|:---:|
| ![temperature_sensor](../../data/equipment/icons/temperature_sensor.svg) | ![humidity_sensor](../../data/equipment/icons/humidity_sensor.svg) | ![ac_unit](../../data/equipment/icons/ac_unit.svg) | ![ventilation_fan](../../data/equipment/icons/ventilation_fan.svg) |
| `temperature_sensor` | `humidity_sensor` | `ac_unit` | `ventilation_fan` |
| ![radiator](../../data/equipment/icons/radiator.svg) | ![thermostat](../../data/equipment/icons/thermostat.svg) | ![ahu](../../data/equipment/icons/ahu.svg) | |
| `radiator` | `thermostat` | `ahu` | |

### Safety & Fire
| | | | |
|:---:|:---:|:---:|:---:|
| ![smoke_detector](../../data/equipment/icons/smoke_detector.svg) | ![fire_extinguisher](../../data/equipment/icons/fire_extinguisher.svg) | ![emergency_light](../../data/equipment/icons/emergency_light.svg) | ![sprinkler](../../data/equipment/icons/sprinkler.svg) |
| `smoke_detector` | `fire_extinguisher` | `emergency_light` | `sprinkler` |
| ![fire_alarm_panel](../../data/equipment/icons/fire_alarm_panel.svg) | ![aed](../../data/equipment/icons/aed.svg) | | |
| `fire_alarm_panel` | `aed` | | |

### Electrical & Power
| | | | |
|:---:|:---:|:---:|:---:|
| ![light_fixture](../../data/equipment/icons/light_fixture.svg) | ![distribution_panel](../../data/equipment/icons/distribution_panel.svg) | ![emergency_generator](../../data/equipment/icons/emergency_generator.svg) | ![ups](../../data/equipment/icons/ups.svg) |
| `light_fixture` | `distribution_panel` | `emergency_generator` | `ups` |

### Plumbing
| | | |
|:---:|:---:|:---:|
| ![water_valve](../../data/equipment/icons/water_valve.svg) | ![water_leak_sensor](../../data/equipment/icons/water_leak_sensor.svg) | ![pump](../../data/equipment/icons/pump.svg) |
| `water_valve` | `water_leak_sensor` | `pump` |

### Monitoring & Sensors
| | | | |
|:---:|:---:|:---:|:---:|
| ![security_camera](../../data/equipment/icons/security_camera.svg) | ![motion_sensor](../../data/equipment/icons/motion_sensor.svg) | ![co2_sensor](../../data/equipment/icons/co2_sensor.svg) | ![co_sensor](../../data/equipment/icons/co_sensor.svg) |
| `security_camera` | `motion_sensor` | `co2_sensor` | `co_sensor` |
| ![air_quality_sensor](../../data/equipment/icons/air_quality_sensor.svg) | ![noise_sensor](../../data/equipment/icons/noise_sensor.svg) | ![light_sensor](../../data/equipment/icons/light_sensor.svg) | ![vibration_sensor](../../data/equipment/icons/vibration_sensor.svg) |
| `air_quality_sensor` | `noise_sensor` | `light_sensor` | `vibration_sensor` |
| ![water_flow_meter](../../data/equipment/icons/water_flow_meter.svg) | ![gas_leak_detector](../../data/equipment/icons/gas_leak_detector.svg) | ![radon_detector](../../data/equipment/icons/radon_detector.svg) | ![occupancy_counter](../../data/equipment/icons/occupancy_counter.svg) |
| `water_flow_meter` | `gas_leak_detector` | `radon_detector` | `occupancy_counter` |

### Network & IT
| | | | | |
|:---:|:---:|:---:|:---:|:---:|
| ![network_switch](../../data/equipment/icons/network_switch.svg) | ![wifi_access_point](../../data/equipment/icons/wifi_access_point.svg) | ![base_station_5g](../../data/equipment/icons/base_station_5g.svg) | ![bms_controller](../../data/equipment/icons/bms_controller.svg) | ![iot_gateway](../../data/equipment/icons/iot_gateway.svg) |
| `network_switch` | `wifi_access_point` | `base_station_5g` | `bms_controller` | `iot_gateway` |

### Vertical Transport
| | |
|:---:|:---:|
| ![elevator](../../data/equipment/icons/elevator.svg) | ![escalator](../../data/equipment/icons/escalator.svg) |
| `elevator` | `escalator` |

### Fixed Equipment
| | | | |
|:---:|:---:|:---:|:---:|
| ![compressor](../../data/equipment/icons/compressor.svg) | ![coffee_machine](../../data/equipment/icons/coffee_machine.svg) | ![printer](../../data/equipment/icons/printer.svg) | ![projector](../../data/equipment/icons/projector.svg) |
| `compressor` | `coffee_machine` | `printer` | `projector` |

### Special (Occupancy & Fallback)
| | | | |
|:---:|:---:|:---:|:---:|
| ![person](../../data/equipment/icons/person.svg) | ![group](../../data/equipment/icons/group.svg) | ![alien](../../data/equipment/icons/alien.svg) | ![generic](../../data/equipment/icons/generic.svg) |
| `person` | `group` | `alien` | `generic` |
