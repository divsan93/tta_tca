#!/bin/bash

host="${host:-localhost:8080}"

# https://github.com/konveyor/tackle-test-generator-cli/blob/main/doc/unit/tkltest_unit_config_options.md
curl -X POST ${host}/tasks -d \
'{
  "name": "Test Generator",
  "state": "Ready",
  "locator": "testgen",
  "addon": "testgen",
  "application": {
    "id": 1
  },
  "data": {
    "branch_name": "ftw",
    "tkltest_config": {
      "generate": {
        "app_build_files": [
          "pom.xml"
        ],
        "target_class_list": [],
        "excluded_class_list": [],
        "time_limit": 30
      }
    }
  }
}' | jq -M .


# curl -X POST ${host}/tasks -d \
# '{
#   "name": "Test Generator",
#   "state": "Ready",
#   "locator": "testgen",
#   "addon": "testgen",
#   "application": {
#     "id": 1
#   },
#   "data": {
#     "branch_name": "ftw",
#     "tkltest_config": {
#       "general": {
#         "test_directory": "tests/"
#       },
#       "generate": {
#         "app_build_files": [
#           "pom.xml"
#         ],
#         "target_class_list": [],
#         "excluded_class_list": [],
#         "time_limit": 30
#       }
#     }
#   }
# }' | jq -M .
