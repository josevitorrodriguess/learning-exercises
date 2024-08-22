import pandas as pd

def selectFirstRows(employees: pd.DataFrame) -> pd.DataFrame:
    data = []

    for i in range(min(3, len(employees))):  
        row = employees.iloc[i]  
        data.append(row)    
   
    result_df = pd.DataFrame(data)
    
    return result_df