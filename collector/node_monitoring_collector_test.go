package collector

import (
	"testing"
)

var JSON1 []byte = []byte(`{
  "host" : "agapiou-Latitude-E6420",
  "version" : "6.4.2",
  "http_address" : "0.0.0.0:9600",
  "id" : "03b69db3-582c-4178-b66c-0d1ef3e60592",
  "name" : "testlogstash",
  "jvm" : {
    "threads" : {
      "count" : 23,
      "peak_count" : 25
    },
    "mem" : {
      "heap_used_percent" : 17,
      "heap_committed_in_bytes" : 1038876672,
      "heap_max_in_bytes" : 1038876672,
      "heap_used_in_bytes" : 184417080,
      "non_heap_used_in_bytes" : 140938168,
      "non_heap_committed_in_bytes" : 161181696,
      "pools" : {
        "survivor" : {
          "peak_used_in_bytes" : 34865152,
          "used_in_bytes" : 31458952,
          "peak_max_in_bytes" : 34865152,
          "max_in_bytes" : 34865152,
          "committed_in_bytes" : 34865152
        },
        "old" : {
          "peak_used_in_bytes" : 127516144,
          "used_in_bytes" : 127516144,
          "peak_max_in_bytes" : 724828160,
          "max_in_bytes" : 724828160,
          "committed_in_bytes" : 724828160
        },
        "young" : {
          "peak_used_in_bytes" : 279183360,
          "used_in_bytes" : 25441984,
          "peak_max_in_bytes" : 279183360,
          "max_in_bytes" : 279183360,
          "committed_in_bytes" : 279183360
        }
      }
    },
    "gc" : {
      "collectors" : {
        "old" : {
          "collection_time_in_millis" : 184,
          "collection_count" : 2
        },
        "young" : {
          "collection_time_in_millis" : 833,
          "collection_count" : 11
        }
      }
    },
    "uptime_in_millis" : 4780041
  },
  "process" : {
    "open_file_descriptors" : 95,
    "peak_open_file_descriptors" : 97,
    "max_file_descriptors" : 1048576,
    "mem" : {
      "total_virtual_in_bytes" : 4788510720
    },
    "cpu" : {
      "total_in_millis" : 239590,
      "percent" : 0,
      "load_average" : {
        "1m" : 0.1,
        "5m" : 0.12,
        "15m" : 0.25
      }
    }
  },
  "events" : {
    "in" : 2,
    "filtered" : 2,
    "out" : 2,
    "duration_in_millis" : 93,
    "queue_push_duration_in_millis" : 0
  },
  "pipelines" : {
    "main" : {
      "events" : {
        "duration_in_millis" : 93,
        "in" : 2,
        "out" : 2,
        "filtered" : 2,
        "queue_push_duration_in_millis" : 0
      },
      "plugins" : {
        "inputs" : [ {
          "id" : "2e751834e57957fe3925003eb2567a0a737fdf74971baefc8cbd483a4a2eee10",
          "events" : {
            "out" : 2,
            "queue_push_duration_in_millis" : 0
          },
          "name" : "stdin"
        } ],
        "filters" : [ ],
        "outputs" : [ {
          "id" : "a9eded44139d380bad7431313723eacb85572dfa2faee0cea1b7504405e917be",
          "events" : {
            "duration_in_millis" : 49,
            "in" : 2,
            "out" : 2
          },
          "name" : "stdout"
        } ]
      },
      "reloads" : {
        "last_error" : null,
        "successes" : 0,
        "last_success_timestamp" : null,
        "last_failure_timestamp" : null,
        "failures" : 0
      },
      "queue" : {
        "type" : "memory"
      },
      "dead_letter_queue" : {
        "queue_size_in_bytes" : 57
      }
    }
  },
  "reloads" : {
    "successes" : 0,
    "failures" : 0
  },
  "os" : {
    "cgroup" : {
      "cpuacct" : {
        "usage_nanos" : 1999652653607,
        "control_group" : "/user.slice"
      },
      "cpu" : {
        "cfs_quota_micros" : -1,
        "control_group" : "/user.slice",
        "stat" : {
          "number_of_times_throttled" : 0,
          "time_throttled_nanos" : 0,
          "number_of_elapsed_periods" : 0
        },
        "cfs_period_micros" : 100000
      }
    }
  }
}`)

func TestNewNodeMonitoringCollector(t *testing.T) {
	var response NodeMonitoringResponse
	m := &MockHTTPHandler{ReturnJSON: JSON1}
	getMetrics(m, &response)
	if response.Events.In != 2 {
		t.Fail()
	}
	if response.Events.Out != 2 {
		t.Fail()
	}
	if response.Events.Filtered != 2 {
		t.Fail()
	}
	if response.Process.CPU.Percent != 0 {
		t.Fail()
	}
	if response.Jvm.Mem.HeapUsedPercent != 17 {
		t.Fail()
	}
}
