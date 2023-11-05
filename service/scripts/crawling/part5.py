test1part5 = driver.find_element(
    by=By.XPATH, value="//*[@id='partcontent-9777']/div[2]"
)

children5 = test1part5.find_elements(By.XPATH, "*")

res = []
for child in children5:
    res.append(child.text)
res[0]


def part5():
    data = {}
    idx = []
    question = []
    answers = []
    nu = []
    for row in res:
        curRow = row.split("\n")
        if len(curRow) == 6:
            idx.append(curRow[0])
            question.append(curRow[1])
            for i in range(4):
                if i != 0:
                    idx.append("")
                    question.append("")
                answers.append(curRow[2 + i])
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

    df.to_csv("part5.csv", index=False)
