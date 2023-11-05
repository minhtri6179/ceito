test1part6 = driver.find_element(
    by=By.XPATH, value="//*[@id='partcontent-9778']/div[2]"
)
children6 = test1part6.find_elements(By.XPATH, "*")

res = []
for block in children6:
    res.append(block.text)


def part6():
    data = {}
    idx = []
    question = []
    answers = []
    nu = []
    for group in res:
        curGroup = group.split("\n")
        curlen = len(curGroup)
        rows = curlen // 5
        for i in range(rows):
            idx.append(curGroup[5 * i])
            for j in range(4):
                if j != 0:
                    idx.append("")
                    question.append("")

                answers.append(curGroup[1 + 5 * i + j])

                nu.append("")
    data = {
        "Question_id": idx,
        "Test_name(ETS23-Test1)": nu,
        "Answer_id": nu,
        "Answer_text": answers,
        "isTrue": nu,
    }
    df = pd.DataFrame(data).reset_index(drop=True)
    df
    df.to_csv("part6.csv", index=False)
