# this file was autogenerated by mysql2postgres.py
# do not edit directly
from sqlalchemy import Column, Integer, String, DateTime, ForeignKey, MetaData
from sqlalchemy import Table
meta = MetaData()


#  ('ca_cloud',)
ca_cloud = Table(
    'ca_cloud', meta,
    Column('wid', Integer ), #int(10) unsigned
    Column('words', String ), #mediumtext
    Column('matrix', String ), #mediumtext
)
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
#  ('ca_document',)
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
#  ('ca_issetobject',)
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
#  ('ca_links',)
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
#  ('ca_metatags',)
ca_metatags = Table(
    'ca_metatags', meta,
    Column('wid', Integer ), #int(10) unsigned
    Column('title', String ), #mediumtext
    Column('keyword', String ), #mediumtext
    Column('description', String ), #mediumtext
    Column('ogproperties', String ), #mediumtext
)
#  ('ca_misc',)
ca_misc = Table(
    'ca_misc', meta,
    Column('wid', Integer ), #int(10) unsigned
    Column('sitemap', String ), #mediumtext
    Column('analytics', String ), #mediumtext
)
#  ('ca_pagespeed',)
ca_pagespeed = Table(
    'ca_pagespeed', meta,
    Column('wid', Integer ), #int(10) unsigned
    Column('data', String ), #longtext
    Column('lang_id', String ), #varchar(5)
)
#  ('ca_w3c',)
ca_w3c = Table(
    'ca_w3c', meta,
    Column('wid', Integer ), #int(10) unsigned
    Column('validator', String ), #enum('html')
    Column('valid', String ), #tinyint(1)
    Column('errors', String ), #smallint(5) unsigned
    Column('warnings', String ), #smallint(5) unsigned
)
#  ('ca_website',)
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
#  ('domain',)
domain = Table(
    'domain', meta,
    Column('id', Integer , primary_key=True), #int(11)
    Column('name', String ), #varchar(64)
    Column('last_crawl_date', String ), #timestamp
)
#  ('feature',)
feature = Table(
    'feature', meta,
    Column('id', Integer , primary_key=True), #int(11)
    Column('name', String ), #varchar(32)
    Column('value', String ), #text
    Column('last_crawl_date', String ), #timestamp
    Column('domain_id', Integer ), #int(11)
    Column('execution_id', Integer ), #int(11)
)
table_list = { 'ca_cloud':ca_cloud,'ca_content':ca_content,'ca_document':ca_document,'ca_issetobject':ca_issetobject,'ca_links':ca_links,'ca_metatags':ca_metatags,'ca_misc':ca_misc,'ca_pagespeed':ca_pagespeed,'ca_w3c':ca_w3c,'ca_website':ca_website,'domain':domain,'feature':feature}
