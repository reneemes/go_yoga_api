# Go Yoga API

YogaFlow is a user-friendly web application designed to create a seamless and engaging experience for yoga enthusiasts. It aims to provide users with the tools to manage their yoga practice and explore additional resources for personal well-being. </br>
This project was initially created and is currently deployed as a Ruby on Rails API. The purpose of this application is to get familiar with Go, Gin, and GORM.

## Links
Rails Repo: https://github.com/reneemes/yoga_flow_be </br>
Deployed Rails API: https://yoga-flow-7a813c31e5f1.herokuapp.com </br>


Angular Front End: https://github.com/reneemes/yoga-flow-fe </br>
Deployed Application: https://yogaflow-app.netlify.app/ </br>

## API Documentation

### GET All Poses
Request: `/api/v1/poses` </br>
Response:
```
{
  "data": [
    {
      "id": 1,
      "name": "Boat",
      "sanskrit_name": "Navasana",
      "translation_name": "nāva = boat, āsana = posture",
      "description": "POSE DESCRIPTION",
      "pose_benefits": "POSE BENEFITS",
      "image_url": "IMG URL"
    },
    {
      "id": 2,
      "name": "Bow",
      "sanskrit_name": "Dhanurasana",
      "translation_name": "dhanur = bow, āsana = posture",
      "description": "POSE DESCRIPTION",
      "pose_benefits": "POSE BENEFITS",
      "image_url": "IMG URL"
    }
  ]
}
```

### GET One Pose
Request: `/api/v1/poses/:id` </br>
Response:
```
{
  "data": {
    "id": 2,
    "type": "pose",
    "attributes": {
        "id": 2,
        "name": "Bow",
        "sanskrit_name": "Dhanurasana",
        "translation_name": "dhanur = bow, āsana = posture",
        "description": "POSE DESCRIPTION",
        "pose_benefits": "POSE BENEFITS",
        "image_url": "IMG URL"
    }
  }
}
```

### GET All Routines
Request: `/api/v1/routines` </br>
Response:
```
{
  "data": [
    {
      "id": 1,
      "name": "Advanced Routine",
      "description": "This is a new Advanced routine.",
      "difficulty": "Advanced",
      "routine_poses": [
        {
          "id": 7,
          "name": "Chair",
          "sanskrit_name": "Utkatasana",
          "translation_name": "utkaṭa = fierce, āsana = posture",
          "description": "POSE DESCRIPTION",
          "pose_benefits": "POSE BENEFITS",
          "image_url": "IMG URL"
        }
      ]
    },
    {
      "id": 2,
      "name": "Advanced Routine",
      "description": "This is a new Advanced routine.",
      "difficulty": "Advanced",
      "routine_poses": [
        {
          "id": 7,
          "name": "Chair",
          "sanskrit_name": "Utkatasana",
          "translation_name": "utkaṭa = fierce, āsana = posture",
          "description": "POSE DESCRIPTION",
          "pose_benefits": "POSE BENEFITS",
          "image_url": "IMG URL"
        }
      ]
    }
  ]
}
```

### GET One Routine
Request: `/api/v1/routines/:id` </br>
Response:
```
{
  "data": {
    "id": 1,
    "name": "Advanced Routine",
    "description": "This is a new advanced yoga routine.",
    "difficulty": "Advanced",
    "routine_poses": [
      {
        "id": 7,
        "name": "Chair",
        "sanskrit_name": "Utkatasana",
        "translation_name": "utkaṭa = fierce, āsana = posture",
        "description": "POSE DESCRIPTION",
        "pose_benefits": "POSE BENEFITS",
        "image_url": "IMG URL"
      },
      {
        "id": 2,
        "name": "Bow",
        "sanskrit_name": "Dhanurasana",
        "translation_name": "dhanur = bow, āsana = posture",
        "description": "POSE DESCRIPTION",
        "pose_benefits": "POSE BENEFITS",
        "image_url": "IMG URL"
      }
    ]
  }
}
```

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```
Create DB container
```bash
make docker-run
```

Shutdown DB Container
```bash
make docker-down
```

DB Integrations Test:
```bash
make itest
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```
