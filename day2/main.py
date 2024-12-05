def parseFile(data):
    with open(data, "r") as file:
        nums = file.read().splitlines()

    reports = []

    for num in nums:
        levels = [int(value) for value in num.split(" ")]

        reports.append(levels)
    return reports


def is_safe_report(report):

    differences = [report[i+1] - report[i] for i in range(len(report) - 1)]
    # print(f"differences: {differences}")

    # Check if all differences are either positive or negative (strictly increasing or decreasing)
    is_increasing = all(diff > 0 for diff in differences)
    is_decreasing = all(diff < 0 for diff in differences)

    if not (is_increasing or is_decreasing):
        return False

    # Check if all differences are between 1 and 3 (inclusive)
    if all(1 <= abs(diff) <= 3 for diff in differences):
        return True

    return False


def count_safe_reports(data):

    safe_count = 0

    for report in data:
        if is_safe_report(report):
            safe_count += 1

    return safe_count


def can_be_safe_by_removing_one(report):

    for i in range(len(report)):
        modified_report = report[:i] + report[i+1:]
        if is_safe_report(modified_report):
            return True

    return False


def count_safe_reports_with_removals(data):

    safe_count = 0

    for report in data:
        if is_safe_report(report) or can_be_safe_by_removing_one(report):
            safe_count += 1

    return safe_count


reports = parseFile("input.txt")
safe_reports = count_safe_reports(reports)
safe_reports_with_removals = count_safe_reports_with_removals(reports)
print(f"Number of safe reports: {safe_reports}")
print(
    f"Number of safe reports including removals: {safe_reports_with_removals}")
