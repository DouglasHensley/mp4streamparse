# mp4streamparse
An MP4 parser with an emphasis on real-time parsing of fragmented MP4 streams.

An MP4 parser is a key component enabling the archive, video-on-demand (VOD) and random access playback functionality in a content distribution server.

### Archive
The initialization fragments - consisting of FTYP and MOOV boxes - are typically stored separate from the video segments, but associated with a range
of MOOF and MDAT pairs for which the initialization is applicable.  
Successive MOOF, MDAT pairs are appended and stored as video segment files. The sum of the MOOF SampleDuration of the aggregated MOOF,MDAT pairs
represent the time duration of each video segment.  
The time length of the video segment represents the granularity that can be supported for VOD playback and streaming modes.

### VOD and Random Access
When a client playback request or slider bar drag/drop random access request is received, the response follows the following steps:

- Identify the applicable initialization fragment for the request timestamp.
- Push this initialization fragment to the client.
- Identify the video segment containing the MOOF, MDAT pairs encompassing the request timestamp.
- Begin pushing video segments to the client from this video segment.

## Install
```
go get -u github.com/DouglasHensley/mp4streamparse
```

The ````examples```` directory contains an example generating a report showing the elements of an MP4 file (````mp4fileparse.go````)
and an example of parsing the elements of an MP4 stream (````mp4streamparse.go````).  



----
### Credit
----
At this time, I am unable to give proper credit to the origin of this project.

In reality, this code is a fork of another GitHub project. The original source was obtained via a ZIP file rather than through git clone or fork.

By the time I was able to push this heavily modified version of the original code for others to access, the provenance was lost to history.  
I apologize to the original author and regret that I benefited from their efforts without giving proper attribution.

----
### License
----
[MIT](http://opensource.org/licenses/MIT)
