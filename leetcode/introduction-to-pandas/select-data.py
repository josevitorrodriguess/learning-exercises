import pandas as pd

def selectData(students: pd.DataFrame) -> pd.DataFrame:
    student = students.loc[students['student_id']==101,['name','age']]

    return student