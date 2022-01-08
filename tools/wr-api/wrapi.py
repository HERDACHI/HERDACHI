# website review api

# thee purpose of this program is to offer a flask api which allows the clients
# to create records in the tables of the database.

# import flask
from flask import Flask, request, jsonify
# import sqlalchemy
from sqlalchemy import create_engine
# import psycopg2
from psycopg2 import sql
# import json
import json
# import os
import os
#from model import *

from sqlalchemy import Column, Integer, String, DateTime, ForeignKey, MetaData
from sqlalchemy import Table
meta = MetaData()


ca_website = Table(
    'ca_website', meta,
    Column('id', Integer , primary_key=True), #int(10) unsigned
    Column('domain', String ), #varchar(255)
    Column('idn', String ), #varchar(255)
    Column('final_url', String ), #mediumtext
    Column('md5domain', String ), #varchar(32)
    Column('added', String ), #timestamp
    Column('modified', String ), #timestamp
    Column('score', String ), #tinyint(3) unsigned
)


app = Flask(__name__)

def get_connection_string():
    return os.getenv("DATABASE_URL")

@app.route('/ca_website/create', methods=['POST'])
def create_ca_website():
    """
    @api {post} /user/:id Create new ca_website
    @apiName CreateCaWebsite
    @apiGroup RecordCreation

    @apiParam {String} domain Domain
    @apiParam {String} idn IDN
    @apiParam {String} final_url Final URL
    @apiParam {String} md5domain MD5 Domain
    @apiParam {String} added Added
    @apiParam {String} modified Modified
    @apiParam {String} score Score

    @apiSuccess {String} message Success message
    
    """
    # get the json object
    json_data = request.get_json()
    # get the data from the json object
    domain = json_data['domain']
    idn = json_data['idn']
    final_url = json_data['final_url']
    md5domain = json_data['md5domain']
    added = json_data['added']
    modified = json_data['modified']
    score = json_data['score']
    # create the sql query
    sql_query = sql.SQL("INSERT INTO ca_website (domain, idn, final_url, md5domain, added, modified, score) VALUES ({}, {}, {}, {}, {}, {}, {})").format(sql.Literal(domain), sql.Literal(idn), sql.Literal(final_url), sql.Literal(md5domain), sql.Literal(added), sql.Literal(modified), sql.Literal(score))
    # connect to the database
    
    engine = create_engine(get_connection_string())
    # execute the query
    engine.execute(sql_query)
    # return the response
    return jsonify({'status': 'success'})


#  ('ca_cloud',)
ca_cloud = Table(
    'ca_cloud', meta,
    Column('wid', Integer ), #int(10) unsigned
    Column('words', String ), #mediumtext
    Column('matrix', String ), #mediumtext
)

@app.route('/ca_cloud/create', methods=['POST'])
def create_ca_cloud():
    """
    @api {post} /user/:id Create new ca_cloud
    @apiName CreateCaCloud
    @apiGroup RecordCreation
    
    @apiParam {Number} wid WID
    @apiParam {String} words Words
    @apiParam {String} matrix Matrix

    @apiSuccess {String} message Success message
    """

    
    # get the json object
    json_data = request.get_json()
    # get the data from the json object
    wid = json_data['wid']
    words = json_data['words']
    matrix = json_data['matrix']
    # create the sql query
    sql_query = sql.SQL("INSERT INTO ca_cloud (wid, words, matrix) VALUES ({}, {}, {})").format(sql.Literal(wid), sql.Literal(words), sql.Literal(matrix))
    # connect to the database
    engine = create_engine(get_connection_string())
    # execute the query
    engine.execute(sql_query)
    # return the response
    return jsonify({'status': 'success'})


#  ('ca_content',)
ca_content = Table(
    'ca_content', meta,
    Column('wid', Integer ), #int(10) unsigned
    Column('headings', String ), #mediumtext
    Column('total_img', Integer ), #int(10) unsigned
    Column('total_alt', Integer ), #int(10) unsigned
    Column('deprecated', String ), #mediumtext
    Column('isset_headings', String ), #tinyint(4)
)
@app.route('/ca_content/create', methods=['POST'])
def create_ca_content():
    """
    @api {post} /user/:id Create new ca_content
    @apiName CreateCaContent
    @apiGroup RecordCreation
    
    @apiParam {Number} wid WID
    @apiParam {String} headings Headings
    @apiParam {Number} total_img Total Img
    @apiParam {Number} total_alt Total Alt
    @apiParam {String} deprecated Deprecated
    @apiParam {String} isset_headings Isset Headings

    @apiSuccess {String} message Success message    
    """
    # get the json object
    json_data = request.get_json()
    # get the data from the json object
    wid = json_data['wid']
    headings = json_data['headings']
    total_img = json_data['total_img']
    total_alt = json_data['total_alt']
    deprecated = json_data['deprecated']
    isset_headings = json_data['isset_headings']
    # create the sql query
    sql_query = sql.SQL("INSERT INTO ca_content (wid, headings, total_img, total_alt, deprecated, isset_headings) VALUES ({}, {}, {}, {}, {}, {})").format(sql.Literal(wid), sql.Literal(headings), sql.Literal(total_img), sql.Literal(total_alt), sql.Literal(deprecated), sql.Literal(isset_headings))
    # connect to the database
    engine = create_engine(get_connection_string())
    # execute the query
    engine.execute(sql_query)
    # return the response
    return jsonify({'status': 'success'})

ca_document = Table(
    'ca_document', meta,
    Column('wid', Integer ), #int(10) unsigned
    Column('doctype', String ), #text
    Column('lang', String ), #varchar(255)
    Column('charset', String ), #varchar(255)
    Column('css', Integer ), #int(10) unsigned
    Column('js', Integer ), #int(10) unsigned
    Column('htmlratio', Integer ), #int(10) unsigned
    Column('favicon', String ), #text
)
@app.route('/ca_document/create', methods=['POST'])
def create_ca_document():
    """
    @api {post} /user/:id Create new ca_document
    @apiName CreateCaDocument
    @apiGroup RecordCreation

    @apiParam {Number} wid WID
    @apiParam {String} doctype Doctype
    @apiParam {String} lang Lang
    @apiParam {String} charset Charset
    @apiParam {Number} css CSS
    @apiParam {Number} js JS
    @apiParam {Number} htmlratio HTML Ratio
    @apiParam {String} favicon Favicon
    
    """
    # get the json object
    json_data = request.get_json()
    # get the data from the json object
    wid = json_data['wid']
    doctype = json_data['doctype']
    lang = json_data['lang']
    charset = json_data['charset']
    css = json_data['css']
    js = json_data['js']
    htmlratio = json_data['htmlratio']
    favicon = json_data['favicon']
    # create the sql query
    sql_query = sql.SQL("INSERT INTO ca_document (wid, doctype, lang, charset, css, js, htmlratio, favicon) VALUES ({}, {}, {}, {}, {}, {}, {}, {})").format(sql.Literal(wid), sql.Literal(doctype), sql.Literal(lang), sql.Literal(charset), sql.Literal(css), sql.Literal(js), sql.Literal(htmlratio), sql.Literal(favicon))
    # connect to the database
    engine = create_engine(get_connection_string()
    # execute the query
    engine.execute(sql_query)
    # return the response
    return jsonify({'status': 'success'})


ca_issetobject = Table(
    'ca_issetobject', meta,
    Column('wid', Integer ), #int(10) unsigned
    Column('flash', String ), #tinyint(1)
    Column('iframe', String ), #tinyint(1)
    Column('nestedtables', String ), #tinyint(1)
    Column('inlinecss', String ), #tinyint(1)
    Column('email', String ), #tinyint(1)
    Column('viewport', String ), #tinyint(1)
    Column('dublincore', String ), #tinyint(1)
    Column('printable', String ), #tinyint(1)
    Column('appleicons', String ), #tinyint(1)
    Column('robotstxt', String ), #tinyint(1)
    Column('gzip', String ), #tinyint(1)
)
@app.route('/ca_issetobject/create', methods=['POST'])
def create_ca_issetobject():
    """
    @api {post} /user/:id Create new ca_issetobject
    @apiName CreateCaIssetobject
    @apiGroup RecordCreation

    @apiParam {Number} wid WID
    @apiParam {String} flash Flash
    @apiParam {String} iframe Iframe
    @apiParam {String} nestedtables Nestedtables
    @apiParam {String} inlinecss Inlinecss
    @apiParam {String} email Email
    @apiParam {String} viewport Viewport
    @apiParam {String} dublincore Dublincore
    @apiParam {String} printable Printable
    @apiParam {String} appleicons Appleicons
    @apiParam {String} robotstxt Robotstxt
    @apiParam {String} gzip Gzip
    
    @apiSuccess {String} message Success message

    """

    
    
    # get the json object
    json_data = request.get_json()
    # get the data from the json object
    wid = json_data['wid']
    flash = json_data['flash']
    iframe = json_data['iframe']
    nestedtables = json_data['nestedtables']
    inlinecss = json_data['inlinecss']
    email = json_data['email']
    viewport = json_data['viewport']
    dublincore = json_data['dublincore']
    printable = json_data['printable']
    appleicons = json_data['appleicons']
    robotstxt = json_data['robotstxt']
    gzip = json_data['gzip']
    # create the sql query
    sql_query = sql.SQL("INSERT INTO ca_issetobject (wid, flash, iframe, nestedtables, inlinecss, email, viewport, dublincore, printable, appleicons, robotstxt, gzip) VALUES ({}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {})").format(sql.Literal(wid), sql.Literal(flash), sql.Literal(iframe), sql.Literal(nestedtables), sql.Literal(inlinecss), sql.Literal(email), sql.Literal(viewport), sql.Literal(dublincore), sql.Literal(printable), sql.Literal(appleicons), sql.Literal(robotstxt), sql.Literal(gzip))
    # connect to the database
    engine = create_engine(get_connection_string())
    # execute the query
    engine.execute(sql_query)
    # return the response
    return jsonify({'status': 'success'})

ca_links = Table(
    'ca_links', meta,
    Column('wid', Integer ), #int(10) unsigned
    Column('links', String ), #mediumtext
    Column('internal', Integer ), #int(10) unsigned
    Column('external_dofollow', Integer ), #int(10) unsigned
    Column('external_nofollow', Integer ), #int(10) unsigned
    Column('isset_underscore', String ), #tinyint(1)
    Column('files_count', Integer ), #int(10) unsigned
    Column('friendly', String ), #tinyint(1)
)
@app.route('/ca_links/create', methods=['POST'])
def create_ca_links():
    """
    @api {post} /user/:id Create new ca_links
    @apiName CreateCaLinks
    @apiGroup RecordCreation
    
    @apiParam {Number} wid WID
    @apiParam {String} links Links
    @apiParam {String} internal Internal
    @apiParam {String} external_dofollow External Dofollow
    @apiParam {String} external_nofollow External Nofollow
    @apiParam {String} isset_underscore Isset Underscore
    @apiParam {String} files_count Files Count
    @apiParam {String} friendly Friendly
    """

    # get the json object
    json_data = request.get_json()
    # get the data from the json object
    wid = json_data['wid']
    links = json_data['links']
    internal = json_data['internal']
    external_dofollow = json_data['external_dofollow']
    external_nofollow = json_data['external_nofollow']
    isset_underscore = json_data['isset_underscore']
    files_count = json_data['files_count']
    friendly = json_data['friendly']
    # create the sql query
    sql_query = sql.SQL("INSERT INTO ca_links (wid, links, internal, external_dofollow, external_nofollow, isset_underscore, files_count, friendly) VALUES ({}, {}, {}, {}, {}, {}, {}, {})").format(sql.Literal(wid), sql.Literal(links), sql.Literal(internal), sql.Literal(external_dofollow), sql.Literal(external_nofollow), sql.Literal(isset_underscore), sql.Literal(files_count), sql.Literal(friendly))
    # connect to the database
    engine = create_engine(get_connection_string())
    # execute the query
    engine.execute(sql_query)
    # return the response
    return jsonify({'status': 'success'})

ca_metatags = Table(
    'ca_metatags', meta,
    Column('wid', Integer ), #int(10) unsigned
    Column('title', String ), #mediumtext
    Column('keyword', String ), #mediumtext
    Column('description', String ), #mediumtext
    Column('ogproperties', String ), #mediumtext
)
@app.route('/ca_metatags/create', methods=['POST'])
def create_ca_metatags():
    """
    @api {post} /user/:id Create new ca_metatags
    @apiName CreateCaMetatags
    @apiGroup RecordCreation
    
    @apiParam {Number} wid WID
    @apiParam {String} title Title
    @apiParam {String} keyword Keyword
    @apiParam {String} description Description
    @apiParam {String} ogproperties Ogproperties
    
    @apiSuccess {String} message Success message

    """

    # get the json object
    json_data = request.get_json()
    # get the data from the json object
    wid = json_data['wid']
    title = json_data['title']
    keyword = json_data['keyword']
    description = json_data['description']
    ogproperties = json_data['ogproperties']
    # create the sql query
    sql_query = sql.SQL("INSERT INTO ca_metatags (wid, title, keyword, description, ogproperties) VALUES ({}, {}, {}, {}, {})").format(sql.Literal(wid), sql.Literal(title), sql.Literal(keyword), sql.Literal(description), sql.Literal(ogproperties))
    # connect to the database
    engine = create_engine(get_connection_string())
    # execute the query
    engine.execute(sql_query)
    # return the response
    return jsonify({'status': 'success'})

ca_misc = Table(
    'ca_misc', meta,
    Column('wid', Integer ), #int(10) unsigned
    Column('sitemap', String ), #mediumtext
    Column('analytics', String ), #mediumtext
)
@app.route('/ca_misc/create', methods=['POST'])
def create_ca_misc():
    """
    @api {post} /user/:id Create new ca_misc
    @apiName CreateCaMisc
    @apiGroup RecordCreation

    @apiParam {Number} wid WID
    @apiParam {String} sitemap Sitemap
    @apiParam {String} analytics Analytics

    @apiSuccess {String} message Success message
    """

    # get the json object
    json_data = request.get_json()
    # get the data from the json object
    wid = json_data['wid']
    sitemap = json_data['sitemap']
    analytics = json_data['analytics']
    # create the sql query
    sql_query = sql.SQL("INSERT INTO ca_misc (wid, sitemap, analytics) VALUES ({}, {}, {})").format(sql.Literal(wid), sql.Literal(sitemap), sql.Literal(analytics))
    # connect to the database
    engine = create_engine(get_connection_string())
    # execute the query
    engine.execute(sql_query)
    # return the response
    return jsonify({'status': 'success'})

ca_pagespeed = Table(
    'ca_pagespeed', meta,
    Column('wid', Integer ), #int(10) unsigned
    Column('data', String ), #longtext
    Column('lang_id', String ), #varchar(5)
)
@app.route('/ca_pagespeed/create', methods=['POST'])
def create_ca_pagespeed():
    """
    @api {post} /user/:id Create new ca_pagespeed
    @apiName CreateCaPagespeed
    @apiGroup RecordCreation

    @apiParam {Number} wid WID
    @apiParam {String} data Data
    @apiParam {String} lang_id Lang ID
    """

    
    # get the json object
    json_data = request.get_json()
    # get the data from the json object
    wid = json_data['wid']
    data = json_data['data']
    lang_id = json_data['lang_id']
    # create the sql query
    sql_query = sql.SQL("INSERT INTO ca_pagespeed (wid, data, lang_id) VALUES ({}, {}, {})").format(sql.Literal(wid), sql.Literal(data), sql.Literal(lang_id))
    # connect to the database
    engine = create_engine(get_connection_string())
    # execute the query
    engine.execute(sql_query)
    # return the response
    return jsonify({'status': 'success'})

ca_w3c = Table(
    'ca_w3c', meta,
    Column('wid', Integer ), #int(10) unsigned
    Column('validator', String ), #enum('html')
    Column('valid', String ), #tinyint(1)
    Column('errors', String ), #smallint(5) unsigned
    Column('warnings', String ), #smallint(5) unsigned
)
@app.route('/ca_w3c/create', methods=['POST'])
def create_ca_w3c():
    """
    @api {post} /user/:id Create new ca_w3c
    @apiName CreateCaW3c
    @apiGroup RecordCreation
    
    @apiParam {Number} wid WID
    @apiParam {String} validator Validator
    @apiParam {String} valid Valid
    @apiParam {String} errors Errors
    @apiParam {String} warnings Warnings
    
    @apiSuccess {String} message Success message
    """
    # get the json object
    json_data = request.get_json()
    # get the data from the json object
    wid = json_data['wid']
    validator = json_data['validator']
    valid = json_data['valid']
    errors = json_data['errors']
    warnings = json_data['warnings']
    # create the sql query
    sql_query = sql.SQL("INSERT INTO ca_w3c (wid, validator, valid, errors, warnings) VALUES ({}, {}, {}, {}, {})").format(sql.Literal(wid), sql.Literal(validator), sql.Literal(valid), sql.Literal(errors), sql.Literal(warnings))
    # connect to the database
    engine = create_engine(get_connection_string())
    # execute the query
    engine.execute(sql_query)
    # return the response
    return jsonify({'status': 'success'})


ca_website = Table(
    'ca_website', meta,
    Column('id', Integer , primary_key=True), #int(10) unsigned
    Column('domain', String ), #varchar(255)
    Column('idn', String ), #varchar(255)
    Column('final_url', String ), #mediumtext
    Column('md5domain', String ), #varchar(32)
    Column('added', String ), #timestamp
    Column('modified', String ), #timestamp
    Column('score', String ), #tinyint(3) unsigned
)
@app.route('/ca_website/create', methods=['POST'])
def create_ca_website():
    """
    @api {post} /user/:id Create new ca_website
    @apiName CreateCaWebsite
    @apiGroup RecordCreation

    @apiParam {Number} id ID
    @apiParam {String} domain Domain
    @apiParam {String} idn IDN
    @apiParam {String} final_url Final URL
    @apiParam {String} md5domain MD5 Domain
    @apiParam {String} added Added
    @apiParam {String} modified Modified
    @apiParam {String} score Score
    """
    
    # get the json object
    json_data = request.get_json()
    # get the data from the json object
    id = json_data['id']
    domain = json_data['domain']
    idn = json_data['idn']
    final_url = json_data['final_url']
    md5domain = json_data['md5domain']
    added = json_data['added']
    modified = json_data['modified']
    score = json_data['score']
    # create the sql query
    sql_query = sql.SQL("INSERT INTO ca_website (id, domain, idn, final_url, md5domain, added, modified, score) VALUES ({}, {}, {}, {}, {}, {}, {}, {})").format(sql.Literal(id), sql.Literal(domain), sql.Literal(idn), sql.Literal(final_url), sql.Literal(md5domain), sql.Literal(added), sql.Literal(modified), sql.Literal(score))
    # connect to the database
    engine = create_engine(get_connection_string()
    # execute the query
    engine.execute(sql_query)
    # return the response
    return jsonify({'status': 'success'})

domain = Table(
    'domain', meta,
    Column('id', Integer , primary_key=True), #int(11)
    Column('name', String ), #varchar(64)
    Column('last_crawl_date', String ), #timestamp
)
@app.route('/domain/create', methods=['POST'])
def create_domain():
    """
    @api {post} /user/:id Create new domain
    @apiName CreateDomain
    @apiGroup RecordCreation
    
    @apiParam {Number} id ID
    @apiParam {String} name Name
    @apiParam {String} last_crawl_date Last crawl date

    @apiSuccess {String} message Success message

    """

    # get the json object
    json_data = request.get_json()
    # get the data from the json object
    id = json_data['id']
    name = json_data['name']
    last_crawl_date = json_data['last_crawl_date']
    # create the sql query
    sql_query = sql.SQL("INSERT INTO domain (id, name, last_crawl_date) VALUES ({}, {}, {})").format(sql.Literal(id), sql.Literal(name), sql.Literal(last_crawl_date))
    # connect to the database
    engine = create_engine(get_connection_string()
    # execute the query
    engine.execute(sql_query)
    # return the response
    return jsonify({'status': 'success'})

feature = Table(
    'feature', meta,
    Column('id', Integer , primary_key=True), #int(11)
    Column('name', String ), #varchar(32)
    Column('value', String ), #text
    Column('last_crawl_date', String ), #timestamp
    Column('domain_id', Integer ), #int(11)
    Column('execution_id', Integer ), #int(11)
)

@app.route('/feature/create', methods=['POST'])
def create_feature():
    """
    @api {post} /user/:id Create new feature
    @apiName CreateFeature
    @apiGroup RecordCreation

    @apiParam {Number} id ID
    @apiParam {String} name Name
    @apiParam {String} value Value
    @apiParam {String} last_crawl_date Last crawl date
    @apiParam {Number} domain_id Domain ID
    @apiParam {Number} execution_id Execution ID

    @apiSuccess {String} message Success message

    
    """
    # get the json object
    json_data = request.get_json()
    # get the data from the json object
    id = json_data['id']
    name = json_data['name']
    value = json_data['value']
    last_crawl_date = json_data['last_crawl_date']
    domain_id = json_data['domain_id']
    execution_id = json_data['execution_id']
    # create the sql query
    sql_query = sql.SQL("INSERT INTO feature (id, name, value, last_crawl_date, domain_id, execution_id) VALUES ({}, {}, {}, {}, {}, {})").format(sql.Literal(id), sql.Literal(name), sql.Literal(value), sql.Literal(last_crawl_date), sql.Literal(domain_id), sql.Literal(execution_id))
    # connect to the database
    engine = create_engine(get_connection_string())
    # execute the query
    engine.execute(sql_query)
    # return the response
    return jsonify({'status': 'success'})

@app.route('/')
def hello_world():
    return 'Server is running'

if __name__ == '__main__':
    app.run(debug=True)
