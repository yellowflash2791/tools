from urllib.parse import urlparse
import urllib.request,ssl,configparser,logging,os.path
from termcolor import colored

from gzip import GzipFile

import re, sys, glob, html, argparse, webbrowser, subprocess, base64, ssl, xml.etree.ElementTree

from string import Template
try:
    from StringIO import StringIO
    readBytesCustom = StringIO
except ImportError:
    from io import BytesIO
    readBytesCustom = BytesIO


def url_to_parse(url):
    '''
    Send requests with Requests
    '''

    q = urllib.request.Request(url)


    q.add_header('User-Agent', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) \
        AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36')
    q.add_header('Accept', 'text/html,\
        application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8')
    q.add_header('Accept-Language', 'en-US,en;q=0.8')
    q.add_header('Accept-Encoding', 'gzip')
#    q.add_header('Cookie', args.cookies)

    try:
        sslcontext = ssl.SSLContext(ssl.PROTOCOL_TLSv1_2)
        response = urllib.request.urlopen(q, context=sslcontext)
    except:
        sslcontext = ssl.SSLContext(ssl.PROTOCOL_TLSv1)
        response = urllib.request.urlopen(q, context=sslcontext)

    if response.info().get('Content-Encoding') == 'gzip':
        data = GzipFile(fileobj=readBytesCustom(response.read())).read()
    elif response.info().get('Content-Encoding') == 'deflate':
        data = response.read().read()
    else:
        data = response.read()
    return data.decode('utf-8', 'replace')


def parsing_url(url,r):
    output=[]
    x=url_to_parse(url)
    regex_name=re.compile(r,re.IGNORECASE)
    match1=regex_name.findall(x)
    for m in match1:
        output.append(list(filter(None, m)))
    return(output)
     
def configuration(u): 
    test=[]
    config = configparser.RawConfigParser()    
    config.read('regex.conf')
    print("Finding information in %s----------------------------------------------------------------------------------------------------------------------------------------------------------------------------"%u)
    parsed_uri = urlparse(u)
    result = '{uri.scheme}://{uri.netloc}'.format(uri=parsed_uri)    
    for c in config.sections():
        r=config[c]['regex']
        off=config[c]['offset']
        col=config[c]['colour']       
        try:
           if c=='ENDPOINTS':
              print(u)
              for x in parsing_url(u,r):
                  for i in x:
                      print(i)
                      if i.startswith('/'):
                         new_url=result+i
                         print(new_url)
                         #url_to_parse(new_url)  
                      else:
                         new_url=result+"/"+i
                         print(new_url)
                         #url_to_parse(new_url)
           else:
           
                  for others in parsing_url(u,r):           
                      print(others) 
   
               
        except urllib.error.HTTPError as e:
             print(e.code)
             return
        except urllib.error.URLError as f:
             print("Use TLS,probably running vulenerable version")
             return
 


def readfile(file):
    f = open(file, "r")
    for u in f.readlines():
        configuration(u)


if __name__ == '__main__':

    try: 

      readfile(str(sys.argv[1]))
      
    except KeyboardInterrupt:
      sys.exit(1)
