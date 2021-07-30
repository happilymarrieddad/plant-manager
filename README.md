Plant Manager
=========================

## Introduction
The plant manager should store all the necessary information about the plants in the garden along with their current status, position, and control of the machinery.

## Data Model Examples
Home Example
```json
[
    "places": [{"id": 1, "name": "My House"}],
    "place_slots": [{"id": 1, "name": "Ground Plot 1", "place_id": 1}],
    "containers": [{"id": 1, "name": "Holes In Ground", "place_slot_id": 1}],
    "container_slots":[
        {"id": 1, "row": 0, "column": 0, "container_id": 1},
        {"id": 2, "row": 0, "column": 1, "container_id": 1},
        {"id": 3, "row": 1, "column": 0, "container_id": 1},
        {"id": 4, "row": 1, "column": 1, "container_id": 1}
    ],
    "plant_types": [{"id": 1, "name": "Blueberries"}],
    "plants": [
        {"id": 1, "container_slot": 3, "plant_type_id": 1}
    ]
] 
```

Facility Example - We use the place as the location
```json
[
    "places": [{"id": 1, "name": "Greenhouse 1"}],
    "place_slots": [{"id": 1, "name": "Machine 1", "place_id": 1}],
    "containers": [{"id": 1, "name": "40ct Plastic Container", "place_slot_id": 2}],
    "container_slots":[
        {"id": 1, "row": 0, "column": 0, "container_id": 1},
        {"id": 2, "row": 0, "column": 1, "container_id": 1},
        {"id": 3, "row": 1, "column": 0, "container_id": 1},
        {"id": 4, "row": 1, "column": 1, "container_id": 1}
    ],
    "plant_types": [{"id": 1, "name": "Blueberries"}],
    "plants": [
        {"id": 1, "container_slot": 3, "plant_type_id": 1}
    ]
] 
```

Facility Example - We use the place a machine at the place (TODO: Add logitude/latitude - city state zip)
```json
[
    "locations": [{"id": 1, "name": "Some Location"}],
    "places": [{"id": 1, "name": "Machine 1", "location_id": 1}],
    "place_slots": [{"id": 1, "name": "Slot 1 of 4 in machine", "place_id": 1}],
    "containers": [{"id": 1, "name": "40ct Plastic Container", "place_slot_id": 2}],
    "container_slots":[
        {"id": 1, "row": 0, "column": 0, "container_id": 1},
        {"id": 2, "row": 0, "column": 1, "container_id": 1},
        {"id": 3, "row": 1, "column": 0, "container_id": 1},
        {"id": 4, "row": 1, "column": 1, "container_id": 1}
    ],
    "plant_types": [{"id": 1, "name": "Blueberries"}],
    "plants": [
        {"id": 1, "container_slot": 3, "plant_type_id": 1}
    ]
] 
```
