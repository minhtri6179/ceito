test1part7 = driver.find_element(
    by=By.XPATH, value="//*[@id='partcontent-9779']/div[2]"
)
children7 = test1part7.find_elements(By.XPATH, "*")
res = []
for block in children7:
    res.append(block.text)


def part7():
    data = {}
    idx = []
    question = []
    answers = []
    nu = []
    for group in res:
        curGroup = group.split("\n")
        curlen = len(curGroup)
        rows = curlen // 6
        for i in range(rows):
            idx.append(curGroup[6 * i])
            question.append(curGroup[6 * i + 1])
            for j in range(4):
                if j != 0:
                    idx.append("")
                    question.append("")

                answers.append(curGroup[2 + 6 * i + j])

                nu.append("")
    data = {
        "Question_id": idx,
        "Question_text": question,
        "Test_name(ETS23-Test1)": nu,
        "Answer_id": nu,
        "Answer_text": answers,
        "isTrue": nu,
    }
    df = pd.DataFrame(data).reset_index(drop=True)
    df
    df.to_csv("part7.csv", index=False)
