#!/usr/bin/env python3

# need port forward metrics-grpc to localhost:50007

import json
import subprocess
from datetime import datetime


def send_grpc_request(queryString: str):
    try:
        # grpcurl -plaintext -proto ./metrics.proto -d <json_string> localhost:50007 Metrics.MetricsService/QueryRange
        result = subprocess.run(
            ["grpcurl", "-plaintext", "-proto", "./metrics.proto",
             "-d", queryString, "localhost:50007", "Metrics.MetricsService/QueryRange"],
            capture_output=True, text=True, check=True
        )
        return result.stdout

    except subprocess.CalledProcessError as e:
        print(f"Error checking {e}")
        raise e


class QueryTimeRange:
    def __init__(self, start, end, step):
        self.start = start
        self.end = end
        self.step = step


def generate_query_string(uuid, app, metric, agg_type, meta, time_range: QueryTimeRange):
    query_obj = {
        "selector": {
            "uuid": uuid,
            "app": app,
            "field": metric,
            "agg_type": agg_type,
            "meta": meta
        },
        "start_sec": time_range.start,
        "end_sec": time_range.end,
        "step_sec": time_range.step
    }

    return json.dumps(query_obj)


def main():
    # now = int(time.time())
    now = 1737279633
    time_range = QueryTimeRange(now - 6 * 24 * 60 * 60, now, 60)

    current_temp_query = generate_query_string(
        "0ba4eddb-0ade-4ddf-b39e-af574b27b88e", "systemp", "current_temp", "value", {}, time_range)

    result = send_grpc_request(current_temp_query)
    result_data = json.loads(result)

    sample_list = result_data["streams"][0]["samples"]

    for i in range(len(sample_list) - 1):
        diff = int(sample_list[i + 1]["timeSec"]) - int(sample_list[i]["timeSec"])
        if (diff >= 600):
            timestamp = int(sample_list[i]["timeSec"])
            date = datetime.fromtimestamp(timestamp)
            print(f"{diff} on time: {date} ({timestamp})")


if __name__ == "__main__":
    # need port forward metric-grpc to localhost:50007
    # need metrics.proto file
    main()
