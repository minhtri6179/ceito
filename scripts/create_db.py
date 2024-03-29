import requests
import pandas as pd
import numpy as np


class CreateTestDB:
    def __init__(self, question_url, answer_url) -> None:
        self.question = question_url
        self.answer = answer_url
        self.questions_idx = []
        self.answers_idx = []

    def create_question(self, question_text_list, test_name):
        question_ids = []
        part = "Part1"
        cur = 0
        for idx, question in enumerate(question_text_list):
            if idx == cur + 6:
                part = "Part2"
            if idx == cur + 6 + 25:
                part = "Part3"
            if idx == cur + 6 + 25 + 39:
                part = "Part4"
            if idx == cur + 6 + 25 + 39 + 30:
                part = "Part5"
            if idx == cur + 6 + 25 + 39 + 30 + 30:
                part = "Part6"
            if idx == cur + 6 + 25 + 39 + 30 + 30 + 16:
                part = "Part7"
            if not question:
                question = "Question"
            myobj = {
                "question_text": question,
                "test_name": test_name + "-" + part,
            }
            x = requests.post(self.question, json=myobj)
            question_ids.append(x.json()["question_id"])
        self.questions_idx = question_ids

    def create_answer(self, answer_text_list, question_id_list, is_correct_list):
        answer_ids = []
        is_correct_list = np.array(is_correct_list, dtype=bool).tolist()
        for idx, answer in enumerate(answer_text_list):
            question_id = question_id_list[idx // 4]
            if not answer:
                answer = "Listening"
            myobj = {
                "question_id": question_id,
                "answer_text": answer,
                "is_correct": is_correct_list[idx],
            }
            print("POST: ", myobj, "to: ", self.answer)
            x = requests.post(self.answer, json=myobj)
            answer_ids.append(x.json()["answer_id"])
        self.answers_idx = answer_ids

    def get_questions(self):
        return self.questions_idx

    def get_answers(self):
        return self.answers_idx

    def update_question(self):
        pass

    def update_answer(self):
        pass


if __name__ == "__main__":
    df = pd.read_csv("ETS23-Test1.csv")
    df = df.where(pd.notnull(df), None)
    question_text = df["Question_text"][::4]
    answer_text = df["Answer_text"]
    no_question = df["Question_id"][::4].astype(int)
    df["is_true"].replace(np.nan, False, inplace=True)
    answer_list = df["is_true"]
    answer_list

    test = CreateTestDB(
        " https://www.ceito.site/questions/", " https://www.ceito.site/answers/"
    )
    test.create_question(question_text, "ETS-23-Test1")
    questions = test.get_questions()

    test.create_answer(answer_text, questions, answer_list)
    answers = test.get_answers()
    print(len(questions))
    print(len(answers))
