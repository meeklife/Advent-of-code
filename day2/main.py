def parsingFIle(filename):

    with open(filename, "r") as file:
        nums = file.read().splitlines()

    reports = []

    for num in nums:
        levels = [int(value) for value in num.split(" ")]
        reports.append(levels)

    # print(f"reports: {reports}")

    for report in reports:

        difference = []

        for i in range(1, len(report)-1):
            diff = abs(report[i] - report[i-1])
            print(diff)

            if diff < 1 or diff > 3:
                print(f"{report}: unsafe ")

            if report[i] > report[i+1]:
                print(f"this {report} is decreasing")
            difference.append(diff)

        # print(f"this {report} has a difference: {difference}")

        # for j in range(len(difference)):
        #     if difference[j] == 1 or difference[j] == 2 or difference[j] == 3:
        #         print(f"this {report} is safe")
        #     else:
        #         print(f"this {report} is not safe")
        #         # if

    return reports


parsingFIle("sample.txt")
