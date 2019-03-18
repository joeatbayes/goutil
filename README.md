# goutil
GoLang Libraries used by my various Go Based Systems such as file2consult,  MDS,  Quantized Classifier, etc.

Contains libraries that I use across several GOLang based programs and wanted to reduce temptation to copy.



## Important Files



## Some other repositories:

* [file2consul](https://github.com/joeatbayes/file2consul)  Utility to load configuration parameters managed in GO into consul or HTTP server.  Supports inheritance,  parameter interpolation and other advanced techniques to minimize manual editing required to support multiple environments. 
* [GoPackaging](https://github.com/joeatbayes/GoPackaging) - Example of how to package a library for direct use from go command line.  Also shows an example program that uses that library.   
* [DevOps Automation with LXD and LXC containers](https://lxddevops.com/) - Scripts to create images that can be booted at will.  Demonstrates layered image creation.  Scripts to launch images,  map ports and to setup layer 5 routing for images ran across many hosts.  Consider an alternative to OpenStack or Kubernetes although it could run on OpenStack.    
* [Quantized Classifier](https://bitbucket.org/joexdobs/ml-classifier-gesture-recognition) -  Advanced machine learning classifier with full examples.  Very fast and competitive for both accuracy and and recall with tensor flow for many problems.  In some instances runs over 50 times faster than tensor flow with comparable precision. 
* [Metadata server MDS](https://bitbucket.org/joexdobs/meta-data-server) A server for storing arbitrary data by ID with very fast retrieval.   Ideal for use with very large data sets when very large scaling is required.   Compare to redis, memcache, memcacheb, riak, but designed to support very high scale on reasonable memory and can handle data sets of many T with good performance.   Pure HTTP based.  Written in GO.   Optimized along line of consuming a queue to keep multiple readers updated in near real-time.   Any single server can guarantee updates but servers as set are still eventual consistency. Measured at supporting over 16K requests at 4K per second  across a set of millions of records sustaining this load for months without degradation.
* [Computer Aided Call Response Engine](https://bitbucket.org/joexdobs/computer-aided-call-response-engine) Supports non technical user definition of call scripts that can be arbitrarily complex scripts. Based on JavaScript to integrate nicely with most RIA applications.  Features script tracking,   Local data storage,  spooling update events to server,   save and restore context for different users to allow call switching. etc.   Simple text based syntax to define the call tree allows rapid deployment and maintenance by non-programmers.
* [Healthcare provider Search](https://bitbucket.org/joexdobs/healthcare-provider-search)  Search UI that provides physician locator functionality.   Scripts to parse and load with over 1.2 million records from CMS.    Allows geo-proximity filtering,  zip to city resolution, etc.   Based on elastic search with node.js middle tier server.  Demonstrates very fast RIA JavaScript which allows hundreds of records faster that many sites render a one.   Take a close look if you want to understand high performance JavaScript.
* [CSV Tables in Browser](https://github.com/joeatbayes/CSVTablesInBrowser) Render CSV files in a browser automatically with little to know code.  Renders them fast and nice looking with automatically repeated headers.   Supports sorting,   Custom formatting by column and can even allow script callback to allow custom data generation for some fields.    Can dramatically reduce work to render large sets of columnar data. 

## License:

Copyright 2018 Joseph Ellsworth  [MIT License](https://opensource.org/licenses/MIT) 

