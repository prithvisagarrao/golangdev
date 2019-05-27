#!/usr/bin/env python
import json
import sys
import time
from datetime import datetime
import calendar
import time
import syslog                                                                                                                                                           
from pprint import pprint
from elasticsearch import Elasticsearch

f = open('/var/ossec/logs/python.log','a+')

args = sys.argv

f.write(repr(args))

src_ip = args[3]
arg_alert_id = args[4]
rule_id = args[5]

time_window = calendar.timegm(time.gmtime()) - 43200

f.write(str(time_window))
es = Elasticsearch([{'host':'10.200.147.78','port':'9200'},{'host':'10.200.147.8                                                                                        0','port':'9200'},{'host':'10.200.147.79','port':'9200'}])


last_login_user_search = {"query":{"bool":{"must":[{"query_string":{"query":"!(d                                                                                        ata.dstuser:*$)","analyze_wildcard":True,"default_field":"*"}},{"match_phrase":{                                                                                        "rule.id":{"query":"18107"}}},{"match_phrase":{"data.id":{"query":"4624"}}},{"ma                                                                                        tch_phrase":{"data.srcip":{"query":src_ip}}},{"range":{"@timestamp":{"gte":time_                                                                                        window,"format":"epoch_second"}}}],"filter":[],"should":[],"must_not":[]}}}

f.write(json.dumps(last_login_user_search))



res = es.search(body=json.dumps(last_login_user_search))

f.write(json.dumps(res))


data_found = res['hits']['hits']

if data_found:
        logon_user = res['hits']['hits'][0]['_source']
        logon_username = logon_user['data']['account_name']
        logon_timestamp = logon_user['@timestamp']
        syslog.syslog("Anomaly detected machine's last login by: "+str(logon_use                                                                                        rname)+" at "+str(logon_timestamp)+".")
else:
        f.write(False)
f.close()
