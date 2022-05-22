# NO pip install spacy-langdetect
# pip install spacy_fastlang

import spacy
from spacy_fastlang import LanguageDetector
from flask import Flask

nlp = spacy.load('en_core_web_sm')
nlp.add_pipe("language_detector")

app = Flask(__name__)

@app.route("/lang_detect/<some_text>")
def lang_detect(some_text):
    doc = nlp(some_text)
    return {"language":doc._.language,"score":doc._.language_score}

