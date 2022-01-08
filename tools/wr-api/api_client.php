<?php
/* the purpose of this class is to behave as a php client to the website review api
 Example usage:

 $client = WebsiteReviewAPIClient()
 $client->create_ca_analytics(array('wid' => '1', 'data' => '', 'lang_id' => ''));
*/

class WebsiteReviewAPIClient {
    function get_base_url(){
        return "http://localhost"
    }

    function make_request($url, $data){

        $options = array(
            'http' => array(
                'header'  => "Content-type: application/json",
                'method'  => 'POST',
                'content' => json_encode($data)
            )
        );
        $context  = stream_context_create($options);
        $result = file_get_contents($url, false, $context);
        if ($result === FALSE) { 
            /* Handle error */ 
            echo "Error calling $url";
        }
        var_dump($result);

    }

    /**
     * 
    #  ('ca_cloud',)
    ca_cloud = Table(
        'ca_cloud', meta,
        Column('wid', Integer ), #int(10) unsigned
        Column('words', String ), #mediumtext
        Column('matrix', String ), #mediumtext
    )
    * 
    * $data = array('wid' => '1', 'words' => '', 'matrix' => '');
    */
    function create_ca_cloud($data){
        $url = get_base_url() . '/ca_cloud/create';
        make_request($url, $data);
    }
    /**
     * #  ('ca_content',)
    ca_content = Table(
        'ca_content', meta,
        Column('wid', Integer ), #int(10) unsigned
        Column('headings', String ), #mediumtext
        Column('total_img', Integer ), #int(10) unsigned
        Column('total_alt', Integer ), #int(10) unsigned
        Column('deprecated', String ), #mediumtext
        Column('isset_headings', String ), #tinyint(4)
    )
    * $data = array('wid' => '1', 'headings' => '', 'total_img' => '', 'total_alt' => '', 'deprecated' => '', 'isset_headings' => '');
    */
    function create_ca_content($data){
        $url = get_base_url() . '/ca_content/create';
        make_request($url, $data);
    }
    /**
     * #  ('ca_document',)
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
        * $data = array('wid' => '1', 'doctype' => '', 'lang' => '', 'charset' => '', 'css' => '', 'js' => '', 'htmlratio' => '', 'favicon' => '');
    */
    function create_ca_document($data){
        $url = get_base_url() . '/ca_document/create';
        make_request($url, $data);
    }

    /**
     * 
    #  ('ca_issetobject',)
     *     ca_issetobject = Table(
     *         'ca_issetobject', meta,
     *         Column('wid', Integer ), #int(10) unsigned
     *         Column('flash', String ), #tinyint(1)
     *         Column('iframe', String ), #tinyint(1)
     *         Column('nestedtables', String ), #tinyint(1)
     *         Column('inlinecss', String ), #tinyint(1)
     *         Column('email', String ), #tinyint(1)
     *         Column('viewport', String ), #tinyint(1)
     *         Column('dublincore', String ), #tinyint(1)
     *         Column('printable', String ), #tinyint(1)
     *         Column('appleicons', String ), #tinyint(1)
     *         Column('robotstxt', String ), #tinyint(1)
     *         Column('gzip', String ), #tinyint(1)
     *     )
     * $data = array('wid' => '1', 'flash' => '', 'iframe' => '', 'nestedtables' => '', 'inlinecss' => '', 'email' => '', 'viewport' => '', 'dublincore' => '', 'printable' => '', 'appleicons' => '', 'robotstxt' => '', 'gzip' => '');
     */
    function create_ca_issetobject($data){
        $url = get_base_url() . '/ca_issetobject/create';
        make_request($url, $data);
    }
    /**
     * ca_links = Table(
     *         'ca_links', meta,
     *         Column('wid', Integer ), #int(10) unsigned
     *         Column('links', String ), #mediumtext
     *         Column('internal', Integer ), #int(10) unsigned
     *         Column('external_dofollow', Integer ), #int(10) unsigned
     *         Column('external_nofollow', Integer ), #int(10) unsigned
     *         Column('isset_underscore', String ), #tinyint(1)
     *         Column('files_count', Integer ), #int(10) unsigned
     *         Column('friendly', String ), #tinyint(1)
     *     )
     * $data = array('wid' => '1', 'links' => '', 'internal' => '', 'external_dofollow' => '', 'external_nofollow' => '', 'isset_underscore' => '', 'files_count' => '', 'friendly' => '');
     */
    function create_ca_links($data){
        $url = get_base_url() . '/ca_links/create';
        make_request($url, $data);
    }
    /**
     * ca_metatags = Table(
     *         'ca_metatags', meta,
     *         Column('wid', Integer ), #int(10) unsigned
     *         Column('title', String ), #mediumtext
     *         Column('keyword', String ), #mediumtext
     *         Column('description', String ), #mediumtext
     *         Column('ogproperties', String ), #mediumtext
     *     )
     * $data = array('wid' => '1', 'title' => '', 'keyword' => '', 'description' => '', 'ogproperties' => '');
     */
    function create_ca_metatags($data){
        $url = get_base_url() . '/ca_metatags/create';
        make_request($url, $data);
    }
    /**
     * #  ('ca_misc',)
     *     ca_misc = Table(
     *         'ca_misc', meta,
     *         Column('wid', Integer ), #int(10) unsigned
     *         Column('sitemap', String ), #mediumtext
     *         Column('analytics', String ), #mediumtext
     *     )
     * $data = array('wid' => '1', 'sitemap' => '', 'analytics' => '');
     */

    function create_ca_misc($data){
        $url = get_base_url() . '/ca_misc/create';
        make_request($url, $data);
    }

    /**
     * ca_pagespeed = Table(
     *         'ca_pagespeed', meta,
     *         Column('wid', Integer ), #int(10) unsigned
     *         Column('data', String ), #longtext
     *         Column('lang_id', String ), #varchar(5)
     *     )
     * $data = array('wid' => '1', 'data' => '', 'lang_id' => '');
     */
    function create_ca_pagespeed($data){
        $url = get_base_url() . '/ca_pagespeed/create';
        make_request($url, $data);
    }

    /**
     * ca_w3c = Table(
     *         'ca_w3c', meta,
     *         Column('wid', Integer ), #int(10) unsigned
     *         Column('validator', String ), #enum('html')
     *         Column('valid', String ), #tinyint(1)
     *         Column('errors', String ), #smallint(5) unsigned
     *         Column('warnings', String ), #smallint(5) unsigned
     *     )
     * $data = array('wid' => '1', 'validator' => 'html', 'valid' => '', 'errors' => '', 'warnings' => '');
     */
    function create_ca_w3c($data){
        $url = get_base_url() . '/ca_w3c/create';
        make_request($url, $data);
    }

    /**
     *     ca_website = Table(
     *         'ca_website', meta,
     *         Column('id', Integer , primary_key=True), #int(10) unsigned
     *         Column('domain', String ), #varchar(255)
     *         Column('idn', String ), #varchar(255)
     *         Column('final_url', String ), #mediumtext
     *         Column('md5domain', String ), #varchar(32)
     *         Column('added', String ), #timestamp
     *         Column('modified', String ), #timestamp
     *         Column('score', String ), #tinyint(3) unsigned
     *     )
     * 
     */
    // Create a new record by using the HTTP API
    // $data = array('domain' => 'www.google.com', 'idn' => '1', 'final_url' => 'www.google.com', 'md5domain' => '', 'added' => '2021-01-01', 'modified' => '2021-01-01', 'score' => '0');
    function create_ca_website($data){
        $url = get_base_url() . '/ca_website/create';
        make_request($url, $data);
    }

    /**
     * #  ('domain',)
     *     domain = Table(
     *         'domain', meta,
     *         Column('id', Integer , primary_key=True), #int(11)
     *         Column('name', String ), #varchar(64)
     *         Column('last_crawl_date', String ), #timestamp
     *     )
     * $data = array('name' => 'www.google.com', 'last_crawl_date' => '2021-01-01');
     */
    function create_domain($data){
        $url = get_base_url() . '/domain/create';
        make_request($url, $data);
    }

    /*
     *     #  ('feature',)
     *     feature = Table(
     *         'feature', meta,
     *         Column('id', Integer , primary_key=True), #int(11)
     *         Column('name', String ), #varchar(32)
     *         Column('value', String ), #text
     *         Column('last_crawl_date', String ), #timestamp
     *         Column('domain_id', Integer ), #int(11)
     *         Column('execution_id', Integer ), #int(11)
     *     )
     * $data = array('name' => '', 'value' => '', 'last_crawl_date' => '2021-01-01', 'domain_id' => '', 'execution_id' => '');
     */
    function create_feature($data){
        $url = get_base_url() . '/feature/create';
        make_request($url, $data);
    }
}
