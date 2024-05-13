import random


def discrete_cdf_with_weights(data, weights):
    total_weight = sum(weights)
    cdf = {}
    cumulative_weight = 0
    for i, value in enumerate(data):
        cumulative_weight += weights[i]
        cdf[value] = cumulative_weight / total_weight
    return cdf


# Example usage:
def weighted_round_robin(resources, weights):
    n = len(resources)
    cdf = discrete_cdf_with_weights(resources, weights)

    random_value = random.random()
    for i in range(n):
        if random_value <= cdf[resources[i]]:
            return resources[i]


# Example usage:

resources = ["A", "B", "C", "D"]
weights = [1, 1, 1, 97]  # Example weights


def run_test(resources, weights):
    schedule_map = {}
    for _ in range(1000):
        scheduled_task = weighted_round_robin(resources, weights)
        if scheduled_task in schedule_map:
            schedule_map[scheduled_task] += 1
        else:
            schedule_map[scheduled_task] = 1

    # Full random experiment
    full_random_schedule_map = {}
    for _ in range(1000):
        n = random.randint(0, len(resources) - 1)
        task = resources[n]
        if task in full_random_schedule_map:
            full_random_schedule_map[task] += 1
        else:
            full_random_schedule_map[task] = 1

    print(
        f"Full schedule map: {full_random_schedule_map}, Weighted with CDF: {schedule_map}"
    )


for _ in range(10):
    run_test(resources, weights)
