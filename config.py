def getConfig(field):
    import json

    f = open('config.json', 'r')
    config = json.load(f)
    f.close()
    
    return config[field]