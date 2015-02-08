package ooxml

import (
	"bytes"
	"encoding/xml"
	"strings"
	"testing"
)

func TestOpenPresentationDocument(t *testing.T) {
	doc, err := OpenPresentationDocument("../testdata/06.Hook into SharePoint APIs with Android.pptx")
	if err != nil {
		t.Fatal(err)
	}
	if doc == nil {
		t.Fatal(err)
	}
	slides := doc.Slides()
	if len(slides) == 0 {
		t.Fatal(doc)
	}
	if txt := doc.String(); strings.TrimSpace(txt) == "" {
		t.Fatal(doc)
	}
	t.Log(doc.String())
}

func TestPresentationDocument_decode(t *testing.T) {
	var buf = bytes.NewBufferString(`
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<p:presentation xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:p="http://schemas.openxmlformats.org/presentationml/2006/main" removePersonalInfoOnSave="1" saveSubsetFonts="1" autoCompressPictures="0">
  <p:sldMasterIdLst><p:sldMasterId id="2147484082" r:id="rId4"/><p:sldMasterId id="2147484046" r:id="rId5"/></p:sldMasterIdLst>
  <p:notesMasterIdLst><p:notesMasterId r:id="rId36"/></p:notesMasterIdLst>
  <p:handoutMasterIdLst><p:handoutMasterId r:id="rId37"/></p:handoutMasterIdLst>
  <p:sldIdLst><p:sldId id="778" r:id="rId6"/><p:sldId id="891" r:id="rId7"/><p:sldId id="780" r:id="rId8"/><p:sldId id="788" r:id="rId9"/><p:sldId id="783" r:id="rId10"/><p:sldId id="913" r:id="rId11"/><p:sldId id="894" r:id="rId12"/><p:sldId id="896" r:id="rId13"/><p:sldId id="895" r:id="rId14"/><p:sldId id="857" r:id="rId15"/><p:sldId id="914" r:id="rId16"/><p:sldId id="911" r:id="rId17"/><p:sldId id="897" r:id="rId18"/><p:sldId id="899" r:id="rId19"/><p:sldId id="901" r:id="rId20"/><p:sldId id="900" r:id="rId21"/><p:sldId id="902" r:id="rId22"/><p:sldId id="903" r:id="rId23"/><p:sldId id="904" r:id="rId24"/><p:sldId id="906" r:id="rId25"/><p:sldId id="858" r:id="rId26"/><p:sldId id="859" r:id="rId27"/><p:sldId id="915" r:id="rId28"/><p:sldId id="912" r:id="rId29"/><p:sldId id="907" r:id="rId30"/><p:sldId id="909" r:id="rId31"/><p:sldId id="910" r:id="rId32"/><p:sldId id="916" r:id="rId33"/><p:sldId id="860" r:id="rId34"/><p:sldId id="654" r:id="rId35"/></p:sldIdLst>
  <p:sldSz cx="12188825" cy="6858000"/><p:notesSz cx="6858000" cy="9144000"/>
  <p:defaultTextStyle><a:defPPr><a:defRPr lang="en-US"/></a:defPPr><a:lvl1pPr marL="0" algn="l" defTabSz="914363" rtl="0" eaLnBrk="1" latinLnBrk="0" hangingPunct="1"><a:defRPr sz="1800" kern="1200"><a:solidFill><a:schemeClr val="tx1"/></a:solidFill><a:latin typeface="+mn-lt"/><a:ea typeface="+mn-ea"/><a:cs typeface="+mn-cs"/></a:defRPr></a:lvl1pPr><a:lvl2pPr marL="457182" algn="l" defTabSz="914363" rtl="0" eaLnBrk="1" latinLnBrk="0" hangingPunct="1"><a:defRPr sz="1800" kern="1200"><a:solidFill><a:schemeClr val="tx1"/></a:solidFill><a:latin typeface="+mn-lt"/><a:ea typeface="+mn-ea"/><a:cs typeface="+mn-cs"/></a:defRPr></a:lvl2pPr><a:lvl3pPr marL="914363" algn="l" defTabSz="914363" rtl="0" eaLnBrk="1" latinLnBrk="0" hangingPunct="1"><a:defRPr sz="1800" kern="1200"><a:solidFill><a:schemeClr val="tx1"/></a:solidFill><a:latin typeface="+mn-lt"/><a:ea typeface="+mn-ea"/><a:cs typeface="+mn-cs"/></a:defRPr></a:lvl3pPr><a:lvl4pPr marL="1371545" algn="l" defTabSz="914363" rtl="0" eaLnBrk="1" latinLnBrk="0" hangingPunct="1"><a:defRPr sz="1800" kern="1200"><a:solidFill><a:schemeClr val="tx1"/></a:solidFill><a:latin typeface="+mn-lt"/><a:ea typeface="+mn-ea"/><a:cs typeface="+mn-cs"/></a:defRPr></a:lvl4pPr><a:lvl5pPr marL="1828727" algn="l" defTabSz="914363" rtl="0" eaLnBrk="1" latinLnBrk="0" hangingPunct="1"><a:defRPr sz="1800" kern="1200"><a:solidFill><a:schemeClr val="tx1"/></a:solidFill><a:latin typeface="+mn-lt"/><a:ea typeface="+mn-ea"/><a:cs typeface="+mn-cs"/></a:defRPr></a:lvl5pPr><a:lvl6pPr marL="2285909" algn="l" defTabSz="914363" rtl="0" eaLnBrk="1" latinLnBrk="0" hangingPunct="1"><a:defRPr sz="1800" kern="1200"><a:solidFill><a:schemeClr val="tx1"/></a:solidFill><a:latin typeface="+mn-lt"/><a:ea typeface="+mn-ea"/><a:cs typeface="+mn-cs"/></a:defRPr></a:lvl6pPr><a:lvl7pPr marL="2743090" algn="l" defTabSz="914363" rtl="0" eaLnBrk="1" latinLnBrk="0" hangingPunct="1"><a:defRPr sz="1800" kern="1200"><a:solidFill><a:schemeClr val="tx1"/></a:solidFill><a:latin typeface="+mn-lt"/><a:ea typeface="+mn-ea"/><a:cs typeface="+mn-cs"/></a:defRPr></a:lvl7pPr><a:lvl8pPr marL="3200272" algn="l" defTabSz="914363" rtl="0" eaLnBrk="1" latinLnBrk="0" hangingPunct="1"><a:defRPr sz="1800" kern="1200"><a:solidFill><a:schemeClr val="tx1"/></a:solidFill><a:latin typeface="+mn-lt"/><a:ea typeface="+mn-ea"/><a:cs typeface="+mn-cs"/></a:defRPr></a:lvl8pPr><a:lvl9pPr marL="3657454" algn="l" defTabSz="914363" rtl="0" eaLnBrk="1" latinLnBrk="0" hangingPunct="1"><a:defRPr sz="1800" kern="1200"><a:solidFill><a:schemeClr val="tx1"/></a:solidFill><a:latin typeface="+mn-lt"/><a:ea typeface="+mn-ea"/><a:cs typeface="+mn-cs"/></a:defRPr></a:lvl9pPr></p:defaultTextStyle>
  <p:extLst>
    <p:ext uri="{521415D9-36F7-43E2-AB2F-B90AF26B5E84}"><p14:sectionLst xmlns:p14="http://schemas.microsoft.com/office/powerpoint/2010/main"><p14:section name="Default Section" id="{53481D9B-1B4A-4A0C-83BE-B54383537EFC}"><p14:sldIdLst><p14:sldId id="778"/><p14:sldId id="891"/><p14:sldId id="780"/><p14:sldId id="788"/></p14:sldIdLst></p14:section><p14:section name="Introduction" id="{9F166A1B-EE90-445C-856F-C8BE0CE6DAFC}"><p14:sldIdLst><p14:sldId id="783"/><p14:sldId id="913"/><p14:sldId id="894"/><p14:sldId id="896"/><p14:sldId id="895"/></p14:sldIdLst></p14:section><p14:section name="Authentication with Azure AD" id="{85741DDF-1794-4EEA-9289-B9F64C100160}"><p14:sldIdLst><p14:sldId id="857"/><p14:sldId id="914"/><p14:sldId id="911"/><p14:sldId id="897"/><p14:sldId id="899"/><p14:sldId id="901"/><p14:sldId id="900"/><p14:sldId id="902"/><p14:sldId id="903"/><p14:sldId id="904"/><p14:sldId id="906"/><p14:sldId id="858"/></p14:sldIdLst></p14:section><p14:section name="O365 SharePoint for Android" id="{8B5F841B-E11B-44EC-A688-549F9FB14ADC}"><p14:sldIdLst><p14:sldId id="859"/><p14:sldId id="915"/><p14:sldId id="912"/><p14:sldId id="907"/><p14:sldId id="909"/><p14:sldId id="910"/><p14:sldId id="916"/><p14:sldId id="860"/><p14:sldId id="654"/></p14:sldIdLst></p14:section></p14:sectionLst></p:ext>
    <p:ext uri="{EFAFB233-063F-42B5-8137-9DF3F51BA10A}"><p15:sldGuideLst xmlns:p15="http://schemas.microsoft.com/office/powerpoint/2012/main"><p15:guide id="1" orient="horz" pos="142"><p15:clr><a:srgbClr val="A4A3A4"/></p15:clr></p15:guide><p15:guide id="2" orient="horz" pos="4176"><p15:clr><a:srgbClr val="A4A3A4"/></p15:clr></p15:guide><p15:guide id="3" orient="horz" pos="912"><p15:clr><a:srgbClr val="A4A3A4"/></p15:clr></p15:guide><p15:guide id="4" orient="horz" pos="1197"><p15:clr><a:srgbClr val="A4A3A4"/></p15:clr></p15:guide><p15:guide id="5" orient="horz" pos="1957"><p15:clr><a:srgbClr val="A4A3A4"/></p15:clr></p15:guide><p15:guide id="6" orient="horz" pos="2723"><p15:clr><a:srgbClr val="A4A3A4"/></p15:clr></p15:guide><p15:guide id="7" orient="horz" pos="2159"><p15:clr><a:srgbClr val="A4A3A4"/></p15:clr></p15:guide><p15:guide id="8" orient="horz" pos="3869"><p15:clr><a:srgbClr val="A4A3A4"/></p15:clr></p15:guide><p15:guide id="9" orient="horz" pos="3572"><p15:clr><a:srgbClr val="A4A3A4"/></p15:clr></p15:guide><p15:guide id="10" pos="128"><p15:clr><a:srgbClr val="A4A3A4"/></p15:clr></p15:guide><p15:guide id="11" pos="1767"><p15:clr><a:srgbClr val="A4A3A4"/></p15:clr></p15:guide><p15:guide id="12" pos="7548"><p15:clr><a:srgbClr val="A4A3A4"/></p15:clr></p15:guide><p15:guide id="13" pos="328"><p15:clr><a:srgbClr val="A4A3A4"/></p15:clr></p15:guide><p15:guide id="14" pos="7353"><p15:clr><a:srgbClr val="A4A3A4"/></p15:clr></p15:guide><p15:guide id="15" pos="613"><p15:clr><a:srgbClr val="A4A3A4"/></p15:clr></p15:guide><p15:guide id="16" pos="7062"><p15:clr><a:srgbClr val="A4A3A4"/></p15:clr></p15:guide><p15:guide id="17" pos="3837"><p15:clr><a:srgbClr val="A4A3A4"/></p15:clr></p15:guide><p15:guide id="18" pos="2216"><p15:clr><a:srgbClr val="A4A3A4"/></p15:clr></p15:guide><p15:guide id="19" pos="3771"><p15:clr><a:srgbClr val="A4A3A4"/></p15:clr></p15:guide></p15:sldGuideLst></p:ext>
    <p:ext uri="{2D200454-40CA-4A62-9FC3-DE9A4176ACB9}"><p15:notesGuideLst xmlns:p15="http://schemas.microsoft.com/office/powerpoint/2012/main"><p15:guide id="1" orient="horz" pos="2880"><p15:clr><a:srgbClr val="A4A3A4"/></p15:clr></p15:guide><p15:guide id="2" pos="2160"><p15:clr><a:srgbClr val="A4A3A4"/></p15:clr></p15:guide></p15:notesGuideLst></p:ext>
  </p:extLst>
</p:presentation>
	`)
	x := &PresentationDocument{}
	dec := xml.NewDecoder(buf)
	if err := dec.Decode(x); err != nil {
		t.Fatal(err)
	}
}

func TestPresentationSlide_decode(t *testing.T) {
	var buf = bytes.NewBufferString(`
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<p:sld xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:p="http://schemas.openxmlformats.org/presentationml/2006/main">
  <p:cSld>
    <p:spTree>
      <p:nvGrpSpPr><p:cNvPr id="1" name=""/><p:cNvGrpSpPr/><p:nvPr/></p:nvGrpSpPr>
      <p:grpSpPr><a:xfrm><a:off x="0" y="0"/><a:ext cx="0" cy="0"/><a:chOff x="0" y="0"/><a:chExt cx="0" cy="0"/></a:xfrm></p:grpSpPr>
      <p:sp><p:nvSpPr><p:cNvPr id="4" name="Subtitle 3"/><p:cNvSpPr><a:spLocks noGrp="1"/></p:cNvSpPr><p:nvPr><p:ph type="subTitle" idx="1"/></p:nvPr></p:nvSpPr><p:spPr/><p:txBody><a:bodyPr/><a:lstStyle/><a:p><a:r><a:rPr lang="en-US" dirty="0" smtClean="0"/><a:t>Authentication with azure AD</a:t></a:r><a:endParaRPr lang="en-US" dirty="0"/></a:p></p:txBody></p:sp>
      <p:sp><p:nvSpPr><p:cNvPr id="2" name="Text Placeholder 1"/><p:cNvSpPr><a:spLocks noGrp="1"/></p:cNvSpPr><p:nvPr><p:ph type="body" sz="quarter" idx="10"/></p:nvPr></p:nvSpPr><p:spPr/><p:txBody><a:bodyPr/><a:lstStyle/><a:p><a:pPr marL="0" indent="0"><a:buNone/></a:pPr><a:r><a:rPr lang="en-US" dirty="0" smtClean="0"/><a:t>demo</a:t></a:r><a:endParaRPr lang="en-US" dirty="0"/></a:p></p:txBody></p:sp>
    </p:spTree>
    <p:extLst>
      <p:ext uri="{BB962C8B-B14F-4D97-AF65-F5344CB8AC3E}"><p14:creationId xmlns:p14="http://schemas.microsoft.com/office/powerpoint/2010/main" val="3536787970"/></p:ext>
    </p:extLst>
  </p:cSld>
  <p:clrMapOvr><a:masterClrMapping/></p:clrMapOvr>
  <p:transition><p:fade/></p:transition>
  <p:timing><p:tnLst><p:par><p:cTn id="1" dur="indefinite" restart="never" nodeType="tmRoot"/></p:par></p:tnLst></p:timing>
</p:sld>
	`)
	x := &PresentationSlide{}
	dec := xml.NewDecoder(buf)
	if err := dec.Decode(x); err != nil {
		t.Fatal(err)
	}
}
