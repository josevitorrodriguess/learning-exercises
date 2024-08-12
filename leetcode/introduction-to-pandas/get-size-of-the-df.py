import pandas as pd

def getDataframeSize(players: pd.DataFrame) -> List[int]:
    
    col,lin = players.shape
    list =[col,lin]

    return list