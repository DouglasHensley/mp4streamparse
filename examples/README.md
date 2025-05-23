# mp4streamparse - Examples

## File Parse Report

### Invocation:
````
go mod init mp4fileparse_example
go mod tidy
go run mp4fileparse.go -f <Input MP4 File Name>
````

### Output Example:  

There are two sections generated.  
The first output section shows the progress of the file parsing and creation of the File Report:  
````
mp4fileparse 2025/04/05 20:39:24.406406 mp4fileparse.go:62: BEGIN: mp4fileparse(catch.mp4, 35701430)
mp4fileparse 2025/04/05 20:39:24.406492 mp4fileparse.go:128: main: Begin Error Group Monitoring
mp4fileparse 2025/04/05 20:39:24.406500 parsefile.go:13: ParseFile: Begin
mp4fileparse 2025/04/05 20:39:24.406518 mp4fileparse.go:70: FileRead: Begin
mp4fileparse 2025/04/05 20:39:24.406571 parsefile.go:54: ParseFile: boxType(ftyp) boxSize(24) workBuff(1316)
mp4fileparse 2025/04/05 20:39:24.406606 parsefile.go:54: ParseFile: boxType(moov) boxSize(756) workBuff(1292)
mp4fileparse 2025/04/05 20:39:24.406707 parsefile.go:54: ParseFile: boxType(moof) boxSize(224) workBuff(484)
mp4fileparse 2025/04/05 20:39:24.406950 parsefile.go:54: ParseFile: boxType(mdat) boxSize(61782) workBuff(62112)
mp4fileparse 2025/04/05 20:39:24.406977 parsefile.go:54: ParseFile: boxType(moof) boxSize(224) workBuff(278)
mp4fileparse 2025/04/05 20:39:24.407215 parsefile.go:54: ParseFile: boxType(mdat) boxSize(61238) workBuff(61906)

. . .  

mp4fileparse 2025/04/05 20:39:27.134450 mp4fileparse.go:104: FileRead: EOF(catch.mp4) Read(35701430)
mp4fileparse 2025/04/05 20:39:27.134460 mp4fileparse.go:115: FileRead: Exiting File Read, Num Reads(27130) BytesRead(35701430)
mp4fileparse 2025/04/05 20:39:27.134468 mp4fileparse.go:117: FileRead: Close(inputMp4Channel)
mp4fileparse 2025/04/05 20:39:27.134476 mp4fileparse.go:119: FileRead: End
mp4fileparse 2025/04/05 20:39:27.134486 parsefile.go:63: ParseFile: Total Bytes Processed(35701430)
````

The second output section is the File Report:  
````
mp4fileparse 2025/04/05 20:39:27.134493 parsefile.go:64: ParseFile: MP4 BoxReport(961):
[
    << FtypBox >>	MajorBrand(iso5) MinorVersion(512) CompatibleBrands([iso6 mp41]) 
    << MoovBox >>
                << MvexBox >>	<nil>
                << MvhdBox >>	Timescale(1000) Duration(0)
                << TrakBox >>
                        << TkhdBox >>	CreationTime(0) ModificationTime(0) TrackID(1) Duration(0)	Layer(0) AlternateGroup(0) Width(83886080) Height(47185920)
                        << MdiaBox >>
                                << HdlrBox >>	Handler(vide) Name(VideoHandler )
                                << MdhdBox >>	CreationTime(0) ModificationTime(0)	Timescale(90000) Duration(0) Language(21956, und)
                                << MinfBox >>
                                        << VmhdBox >>	GraphicsMode(0) OpColor(0)
                                        << StblBox >>
                                                << SttsBox >>	EntryCount(0)	SampleCounts([])	SampleDeltas([])
                                                << StsdBox >>	EntryCount(1)	SampleEntry([{0}])	SampleDeltas(                            << Avc1Box >>	DataRefIndex(1)  Width(1280) Height(720) H-Resolution(72.000000) V-Resolution(72.000000)
                                <nil>
                                <nil>)
                    <nil> 
    << MoofBox >>	PrevSeqNo(0) ElapsedTime(0.271989)
                << MfhdBox >>	SequenceNumber(1)
                << TrafBox >>
                        << TfhdBox >>	TrackID(1) BaseDataOffset(0) SampleDescriptionIndex(0) Default Sample: Duration(1727) Size(41463) Flags(16842752)
                        << TfdtBox >>	BaseMediaDecodeTime(0)
                        << TrunBox >>	SampleCount(15) DataOffset(0)	SampleArray([<< TrunSample >> Duration(1727) Size(41463) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1718) Size(1033) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1483) Size(972) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1563) Size(1097) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1689) Size(1596) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1946) Size(1535) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1952) Size(1212) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1711) Size(1497) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1659) Size(1696) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1406) Size(1645) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1595) Size(1510) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1421) Size(1620) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1529) Size(1726) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1550) Size(1686) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1530) Size(1486) Flags(0) CompositionTimeOffset(0)]) 
    << MdatBox >>	Data Length(61782) 
    << MoofBox >>	PrevSeqNo(0) ElapsedTime(0.562756)
                << MfhdBox >>	SequenceNumber(2)
                << TrafBox >>
                        << TfhdBox >>	TrackID(1) BaseDataOffset(0) SampleDescriptionIndex(0) Default Sample: Duration(1754) Size(41536) Flags(16842752)
                        << TfdtBox >>	BaseMediaDecodeTime(0)
                        << TrunBox >>	SampleCount(15) DataOffset(0)	SampleArray([<< TrunSample >> Duration(1754) Size(41536) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1825) Size(635) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1931) Size(1261) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(2166) Size(1370) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1809) Size(1268) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1546) Size(1568) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1613) Size(1788) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1492) Size(1192) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1434) Size(462) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1657) Size(1942) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1744) Size(1492) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1686) Size(1584) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1823) Size(1570) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1900) Size(1800) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1789) Size(1762) Flags(0) CompositionTimeOffset(0)]) 
    << MdatBox >>	Data Length(61238) 
````

## Stream Parse

### Invocation:
````
go mod init mp4streamparse_example
go mod tidy
go run mp4streamparse.go -f <Input MP4 File Name>
````

### Output Example:  

````
mp4streamparse 2025/04/06 16:09:52.590215 mp4streamparse.go:65: >>>>>>>>>> BEGIN <<<<<<<<<<
mp4streamparse 2025/04/06 16:09:52.590326 mp4streamparse.go:154: main: Begin Error Group Monitoring
mp4streamparse 2025/04/06 16:09:52.590349 mp4streamparse.go:95: FileRead: Begin
mp4streamparse 2025/04/06 16:09:52.590372 parsestream.go:16: ParseStream: Begin
mp4streamparse 2025/04/06 16:09:52.613627 mp4streamparse.go:87: fnTsBoxStream: 
## TimestampBox ##
    << FtypBox >>	MajorBrand(iso5) MinorVersion(512) CompatibleBrands([iso6 mp41])
mp4streamparse 2025/04/06 16:09:52.613712 mp4streamparse.go:87: fnTsBoxStream: 
## TimestampBox ##
    << MoovBox >>
                << MvexBox >>	<nil>
                << MvhdBox >>	Timescale(1000) Duration(0)
                << TrakBox >>
                        << TkhdBox >>	CreationTime(0) ModificationTime(0) TrackID(1) Duration(0)	Layer(0) AlternateGroup(0) Width(83886080) Height(47185920)
                        << MdiaBox >>
                                << HdlrBox >>	Handler(vide) Name(VideoHandler )
                                << MdhdBox >>	CreationTime(0) ModificationTime(0)	Timescale(90000) Duration(0) Language(21956, und)
                                << MinfBox >>
                                        << VmhdBox >>	GraphicsMode(0) OpColor(0)
                                        << StblBox >>
                                                << SttsBox >>	EntryCount(0)	SampleCounts([])	SampleDeltas([])
                                                << StsdBox >>	EntryCount(1)	SampleEntry([{0}])	SampleDeltas(                            << Avc1Box >>	DataRefIndex(1)  Width(1280) Height(720) H-Resolution(72.000000) V-Resolution(72.000000)
                                <nil>
                                <nil>)
                    <nil>
mp4streamparse 2025/04/06 16:09:52.613864 mp4streamparse.go:87: fnTsBoxStream: 
## TimestampBox ##
    << MoofBox >>	PrevSeqNo(0) ElapsedTime(0.271989)
                << MfhdBox >>	SequenceNumber(1)
                << TrafBox >>
                        << TfhdBox >>	TrackID(1) BaseDataOffset(0) SampleDescriptionIndex(0) Default Sample: Duration(1727) Size(41463) Flags(16842752)
                        << TfdtBox >>	BaseMediaDecodeTime(0)
                        << TrunBox >>	SampleCount(15) DataOffset(0)	SampleArray([<< TrunSample >> Duration(1727) Size(41463) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1718) Size(1033) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1483) Size(972) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1563) Size(1097) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1689) Size(1596) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1946) Size(1535) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1952) Size(1212) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1711) Size(1497) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1659) Size(1696) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1406) Size(1645) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1595) Size(1510) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1421) Size(1620) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1529) Size(1726) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1550) Size(1686) Flags(0) CompositionTimeOffset(0) << TrunSample >> Duration(1530) Size(1486) Flags(0) CompositionTimeOffset(0)])
mp4streamparse 2025/04/06 16:09:52.614738 mp4streamparse.go:87: fnTsBoxStream: 
## TimestampBox ##
    << MdatBox >>	Data Length(61782)
````