package main

import (
	"strings"
	"testing"
)

func TestInfoRequest(t *testing.T) {
	got, err := carregaItinerario("5566")
	if err != nil || false == strings.Contains(string(got), "VILA NOVA") {
		t.Fatal("Falhei", err, string(got))
	}
}

func TestItinerarioPayload(t *testing.T) {
	//req, err := infoRequest("5566")
	req := `{"idlinha":"5566","nome":"VILA NOVA","codigo":"266-1","0":{"lat":"-30.12419057422600000","lng":"-51.22378313620700000"},"1":{"lat":"-30.12410057422600000","lng":"-51.22352313620700000"},"2":{"lat":"-30.12373357422600000","lng":"-51.22265713620700000"},"3":{"lat":"-30.12305757422600000","lng":"-51.22116713620700000"},"4":{"lat":"-30.12301857422600000","lng":"-51.22119413620700000"},"5":{"lat":"-30.12262857422600000","lng":"-51.22032813620700000"},"6":{"lat":"-30.12223457422600000","lng":"-51.21949313620700000"},"7":{"lat":"-30.12161657422600000","lng":"-51.21815113620700000"},"8":{"lat":"-30.12123957422600000","lng":"-51.21735113620700000"},"9":{"lat":"-30.12094357422600000","lng":"-51.21668113620700000"},"10":{"lat":"-30.12087657422600000","lng":"-51.21652913620700000"},"11":{"lat":"-30.12084357422600000","lng":"-51.21638313620700000"},"12":{"lat":"-30.12069857422600000","lng":"-51.21590913620700000"},"13":{"lat":"-30.12061057422600000","lng":"-51.21563713620700000"},"14":{"lat":"-30.12053257422600000","lng":"-51.21543113620700000"},"15":{"lat":"-30.12026157422600000","lng":"-51.21489813620700000"},"16":{"lat":"-30.11971057422600000","lng":"-51.21384913620700000"},"17":{"lat":"-30.11941857422600000","lng":"-51.21302913620700000"},"18":{"lat":"-30.11869657422600000","lng":"-51.21113913620700000"},"19":{"lat":"-30.11771957422600000","lng":"-51.20740013620700000"},"20":{"lat":"-30.11767257422600000","lng":"-51.20709413620700000"},"21":{"lat":"-30.11775157422600000","lng":"-51.20656713620700000"},"22":{"lat":"-30.11776757422600000","lng":"-51.20644713620700000"},"23":{"lat":"-30.11775557422600000","lng":"-51.20627713620700000"},"24":{"lat":"-30.11772857422600000","lng":"-51.20621313620700000"},"25":{"lat":"-30.11764257422600000","lng":"-51.20618413620700000"},"26":{"lat":"-30.11755057422600000","lng":"-51.20624113620700000"},"27":{"lat":"-30.11750757422600000","lng":"-51.20633113620700000"},"28":{"lat":"-30.11750157422600000","lng":"-51.20647213620700000"},"29":{"lat":"-30.11751857422600000","lng":"-51.20677313620700000"},"30":{"lat":"-30.11744757422600000","lng":"-51.20705413620700000"},"31":{"lat":"-30.11734657422600000","lng":"-51.20726913620700000"},"32":{"lat":"-30.11713657422600000","lng":"-51.20765613620700000"},"33":{"lat":"-30.11653457422600000","lng":"-51.20866613620700000"},"34":{"lat":"-30.11631857422600000","lng":"-51.20915513620700000"},"35":{"lat":"-30.11620657422600000","lng":"-51.20945113620700000"},"36":{"lat":"-30.11596257422600000","lng":"-51.21033113620700000"},"37":{"lat":"-30.11577757422600000","lng":"-51.21113913620700000"},"38":{"lat":"-30.11563857422600000","lng":"-51.21186713620700000"},"39":{"lat":"-30.11550957422600000","lng":"-51.21246413620700000"},"40":{"lat":"-30.11537557422600000","lng":"-51.21306013620700000"},"41":{"lat":"-30.11520957422600000","lng":"-51.21385213620700000"},"42":{"lat":"-30.11512357422600000","lng":"-51.21405013620700000"},"43":{"lat":"-30.11498557422600000","lng":"-51.21423613620700000"},"44":{"lat":"-30.11463657422600000","lng":"-51.21464013620700000"},"45":{"lat":"-30.11451157422600000","lng":"-51.21480913620700000"},"46":{"lat":"-30.11442157422600000","lng":"-51.21497013620700000"},"47":{"lat":"-30.11429557422600000","lng":"-51.21526813620700000"},"48":{"lat":"-30.11416157422600000","lng":"-51.21568913620700000"},"49":{"lat":"-30.11396057422600000","lng":"-51.21656013620700000"},"50":{"lat":"-30.11386857422600000","lng":"-51.21695613620700000"},"51":{"lat":"-30.11364357422600000","lng":"-51.21801213620700000"},"52":{"lat":"-30.11333657422600000","lng":"-51.21937513620700000"},"53":{"lat":"-30.11314057422600000","lng":"-51.22027913620700000"},"54":{"lat":"-30.11275257422600000","lng":"-51.22198513620700000"},"55":{"lat":"-30.11274257422600000","lng":"-51.22288813620700000"},"56":{"lat":"-30.11277457422600000","lng":"-51.22341913620700000"},"57":{"lat":"-30.11285457422600000","lng":"-51.22396713620700000"},"58":{"lat":"-30.11301257422600000","lng":"-51.22475313620700000"},"59":{"lat":"-30.11305357422600000","lng":"-51.22507913620700000"},"60":{"lat":"-30.11305857422600000","lng":"-51.22554813620700000"},"61":{"lat":"-30.11294257422600000","lng":"-51.22635113620700000"},"62":{"lat":"-30.11293857422600000","lng":"-51.22657413620700000"},"63":{"lat":"-30.11217657422600000","lng":"-51.22671713620700000"},"64":{"lat":"-30.11058957422600000","lng":"-51.22696713620700000"},"65":{"lat":"-30.10946157422600000","lng":"-51.22707113620700000"},"66":{"lat":"-30.10909357422600000","lng":"-51.22710313620700000"},"67":{"lat":"-30.10875957422600000","lng":"-51.22715613620700000"},"68":{"lat":"-30.10829457422600000","lng":"-51.22724513620700000"},"69":{"lat":"-30.10794757422600000","lng":"-51.22734713620700000"},"70":{"lat":"-30.10772757422600000","lng":"-51.22745213620700000"},"71":{"lat":"-30.10763057422600000","lng":"-51.22751513620700000"},"72":{"lat":"-30.10746857422600000","lng":"-51.22764113620700000"},"73":{"lat":"-30.10618657422600000","lng":"-51.22925413620700000"},"74":{"lat":"-30.10589757422600000","lng":"-51.22949713620700000"},"75":{"lat":"-30.10565457422600000","lng":"-51.22959513620700000"},"76":{"lat":"-30.10452157422600000","lng":"-51.23002313620700000"},"77":{"lat":"-30.10437157422600000","lng":"-51.23007913620700000"},"78":{"lat":"-30.10362257422600000","lng":"-51.23034713620700000"},"79":{"lat":"-30.10319957422600000","lng":"-51.23055613620700000"},"80":{"lat":"-30.10266257422600000","lng":"-51.23088113620700000"},"81":{"lat":"-30.10237657422600000","lng":"-51.23105813620700000"},"82":{"lat":"-30.10205657422600000","lng":"-51.23107213620700000"},"83":{"lat":"-30.09955057422600000","lng":"-51.23144313620700000"},"84":{"lat":"-30.09865957422600000","lng":"-51.23156913620700000"},"85":{"lat":"-30.09855357422600000","lng":"-51.23158313620700000"},"86":{"lat":"-30.09805757422600000","lng":"-51.23165313620700000"},"87":{"lat":"-30.09788257422600000","lng":"-51.23161613620700000"},"88":{"lat":"-30.09773257422600000","lng":"-51.23153613620700000"},"89":{"lat":"-30.09765757422600000","lng":"-51.23147213620700000"},"90":{"lat":"-30.09756357422600000","lng":"-51.23136413620700000"},"91":{"lat":"-30.09748657422600000","lng":"-51.23122513620700000"},"92":{"lat":"-30.09743457422600000","lng":"-51.23106313620700000"},"93":{"lat":"-30.09741257422600000","lng":"-51.23088713620700000"},"94":{"lat":"-30.09733057422600000","lng":"-51.23013913620700000"},"95":{"lat":"-30.09719957422600000","lng":"-51.22951513620700000"},"96":{"lat":"-30.09709157422600000","lng":"-51.22899813620700000"},"97":{"lat":"-30.09690957422600000","lng":"-51.22820413620700000"},"98":{"lat":"-30.09685957422600000","lng":"-51.22803513620700000"},"99":{"lat":"-30.09678157422600000","lng":"-51.22784713620700000"},"100":{"lat":"-30.09656357422600000","lng":"-51.22746813620700000"},"101":{"lat":"-30.09649457422600000","lng":"-51.22736413620700000"},"102":{"lat":"-30.09640757422600000","lng":"-51.22723913620700000"},"103":{"lat":"-30.09583657422600000","lng":"-51.22656313620700000"},"104":{"lat":"-30.09524757422600000","lng":"-51.22587613620700000"},"105":{"lat":"-30.09515557422600000","lng":"-51.22574113620700000"},"106":{"lat":"-30.09504357422600000","lng":"-51.22550713620700000"},"107":{"lat":"-30.09484857422600000","lng":"-51.22515913620700000"},"108":{"lat":"-30.09467657422600000","lng":"-51.22490813620700000"},"109":{"lat":"-30.09446857422600000","lng":"-51.22469413620700000"},"110":{"lat":"-30.09343857422600000","lng":"-51.22375413620700000"},"111":{"lat":"-30.09322057422600000","lng":"-51.22353913620700000"},"112":{"lat":"-30.09229157422600000","lng":"-51.22239413620700000"},"113":{"lat":"-30.09223857422600000","lng":"-51.22232913620700000"},"114":{"lat":"-30.09193857422600000","lng":"-51.22195313620700000"},"115":{"lat":"-30.09177557422600000","lng":"-51.22164213620700000"},"116":{"lat":"-30.09148557422600000","lng":"-51.22087713620700000"},"117":{"lat":"-30.09125357422600000","lng":"-51.22028513620700000"},"118":{"lat":"-30.09118457422600000","lng":"-51.22007913620700000"},"119":{"lat":"-30.09115957422600000","lng":"-51.22005213620700000"},"120":{"lat":"-30.09084057422600000","lng":"-51.21971113620700000"},"121":{"lat":"-30.08965857422600000","lng":"-51.21857513620700000"},"122":{"lat":"-30.08862057422600000","lng":"-51.21758513620700000"},"123":{"lat":"-30.08746257422600000","lng":"-51.21647213620700000"},"124":{"lat":"-30.08737157422600000","lng":"-51.21640213620700000"},"125":{"lat":"-30.08677457422600000","lng":"-51.21596713620700000"},"126":{"lat":"-30.08598457422600000","lng":"-51.21542013620700000"},"127":{"lat":"-30.08483957422600000","lng":"-51.21464713620700000"},"128":{"lat":"-30.08400057422600000","lng":"-51.21405613620700000"},"129":{"lat":"-30.08379857422600000","lng":"-51.21392513620700000"},"130":{"lat":"-30.08363157422600000","lng":"-51.21382113620700000"},"131":{"lat":"-30.08325657422600000","lng":"-51.21358613620700000"},"132":{"lat":"-30.08248757422600000","lng":"-51.21331213620700000"},"133":{"lat":"-30.08237057422600000","lng":"-51.21318113620700000"},"134":{"lat":"-30.08221657422600000","lng":"-51.21291513620700000"},"135":{"lat":"-30.08197557422600000","lng":"-51.21253213620700000"},"136":{"lat":"-30.08185957422600000","lng":"-51.21237313620700000"},"137":{"lat":"-30.08160357422600000","lng":"-51.21193413620700000"},"138":{"lat":"-30.08145157422600000","lng":"-51.21161813620700000"},"139":{"lat":"-30.08140357422600000","lng":"-51.21148013620700000"},"140":{"lat":"-30.08134857422600000","lng":"-51.21124013620700000"},"141":{"lat":"-30.08121857422600000","lng":"-51.21058813620700000"},"142":{"lat":"-30.08085557422600000","lng":"-51.21026513620700000"},"143":{"lat":"-30.08077557422600000","lng":"-51.21019913620700000"},"144":{"lat":"-30.08054757422600000","lng":"-51.21007813620700000"},"145":{"lat":"-30.07974557422600000","lng":"-51.20948413620700000"},"146":{"lat":"-30.07914157422600000","lng":"-51.20905513620700000"},"147":{"lat":"-30.07825157422600000","lng":"-51.20843313620700000"},"148":{"lat":"-30.07791557422600000","lng":"-51.20831213620700000"},"149":{"lat":"-30.07711957422600000","lng":"-51.20803513620700000"},"150":{"lat":"-30.07522857422600000","lng":"-51.20809613620700000"},"151":{"lat":"-30.07513357422600000","lng":"-51.20811813620700000"},"152":{"lat":"-30.07493957422600000","lng":"-51.20825613620700000"},"153":{"lat":"-30.07470757422600000","lng":"-51.20842213620700000"},"154":{"lat":"-30.07431057422600000","lng":"-51.20863513620700000"},"155":{"lat":"-30.07359657422600000","lng":"-51.20897513620700000"},"156":{"lat":"-30.07343557422600000","lng":"-51.20901813620700000"},"157":{"lat":"-30.07325657422600000","lng":"-51.20902013620700000"},"158":{"lat":"-30.07310157422600000","lng":"-51.20899513620700000"},"159":{"lat":"-30.07255857422600000","lng":"-51.20879813620700000"},"160":{"lat":"-30.07213957422600000","lng":"-51.20866513620700000"},"161":{"lat":"-30.07191257422600000","lng":"-51.20863913620700000"},"162":{"lat":"-30.07164557422600000","lng":"-51.20872713620700000"},"163":{"lat":"-30.07155257422600000","lng":"-51.20883013620700000"},"164":{"lat":"-30.07145357422600000","lng":"-51.20928613620700000"},"165":{"lat":"-30.07139857422600000","lng":"-51.20950013620700000"},"166":{"lat":"-30.07142157422600000","lng":"-51.20954213620700000"},"167":{"lat":"-30.07132557422600000","lng":"-51.21007813620700000"},"168":{"lat":"-30.07125557422600000","lng":"-51.21037313620700000"},"169":{"lat":"-30.07116257422600000","lng":"-51.21057313620700000"},"170":{"lat":"-30.07107557422600000","lng":"-51.21069113620700000"},"171":{"lat":"-30.07020657422600000","lng":"-51.21047713620700000"},"172":{"lat":"-30.06908557422600000","lng":"-51.21025013620700000"},"173":{"lat":"-30.06892857422600000","lng":"-51.21099913620700000"},"174":{"lat":"-30.06818157422600000","lng":"-51.21086313620700000"},"175":{"lat":"-30.06795057422600000","lng":"-51.21082213620700000"},"176":{"lat":"-30.06698257422600000","lng":"-51.21065013620700000"},"177":{"lat":"-30.06681957422600000","lng":"-51.21062113620700000"},"178":{"lat":"-30.06598257422600000","lng":"-51.21047913620700000"},"179":{"lat":"-30.06548757422600000","lng":"-51.21243713620700000"},"180":{"lat":"-30.06449857422600000","lng":"-51.21225113620700000"},"181":{"lat":"-30.06322157422600000","lng":"-51.21200113620700000"},"182":{"lat":"-30.06246457422600000","lng":"-51.21184213620700000"},"183":{"lat":"-30.06175557422600000","lng":"-51.21169513620700000"},"184":{"lat":"-30.06108957422600000","lng":"-51.21155913620700000"},"185":{"lat":"-30.06060557422600000","lng":"-51.21159713620700000"},"186":{"lat":"-30.06053357422600000","lng":"-51.21166513620700000"},"187":{"lat":"-30.05988057422600000","lng":"-51.21214713620700000"},"188":{"lat":"-30.05954957422600000","lng":"-51.21237513620700000"},"189":{"lat":"-30.05904657422600000","lng":"-51.21274813620700000"},"190":{"lat":"-30.05820757422600000","lng":"-51.21331713620700000"},"191":{"lat":"-30.05800857422600000","lng":"-51.21345913620700000"},"192":{"lat":"-30.05597757422600000","lng":"-51.21495613620700000"},"193":{"lat":"-30.05591057422600000","lng":"-51.21492813620700000"},"194":{"lat":"-30.05577557422600000","lng":"-51.21486913620700000"},"195":{"lat":"-30.05515457422600000","lng":"-51.21467913620700000"},"196":{"lat":"-30.05402457422600000","lng":"-51.21422713620700000"},"197":{"lat":"-30.05286757422600000","lng":"-51.21374713620700000"},"198":{"lat":"-30.05179457422600000","lng":"-51.21328913620700000"},"199":{"lat":"-30.05113157422600000","lng":"-51.21301413620700000"},"200":{"lat":"-30.05106857422600000","lng":"-51.21290113620700000"},"201":{"lat":"-30.05066157422600000","lng":"-51.21220213620700000"},"202":{"lat":"-30.04989157422600000","lng":"-51.21088713620700000"},"203":{"lat":"-30.04938657422600000","lng":"-51.21003813620700000"},"204":{"lat":"-30.04856757422600000","lng":"-51.21068513620700000"},"205":{"lat":"-30.04777657422600000","lng":"-51.21130113620700000"},"206":{"lat":"-30.04716257422600000","lng":"-51.21177413620700000"},"207":{"lat":"-30.04693757422600000","lng":"-51.21194713620700000"},"208":{"lat":"-30.04673957422600000","lng":"-51.21209913620700000"},"209":{"lat":"-30.04594357422600000","lng":"-51.21271713620700000"},"210":{"lat":"-30.04527257422600000","lng":"-51.21323013620700000"},"211":{"lat":"-30.04436757422600000","lng":"-51.21392813620700000"},"212":{"lat":"-30.04379057422600000","lng":"-51.21442913620700000"},"213":{"lat":"-30.04347057422600000","lng":"-51.21466513620700000"},"214":{"lat":"-30.04282557422600000","lng":"-51.21514313620700000"},"215":{"lat":"-30.04217357422600000","lng":"-51.21563613620700000"},"216":{"lat":"-30.04117857422600000","lng":"-51.21641013620700000"},"217":{"lat":"-30.04027357422600000","lng":"-51.21709913620700000"},"218":{"lat":"-30.03990657422600000","lng":"-51.21738613620700000"},"219":{"lat":"-30.03923557422600000","lng":"-51.21790913620700000"},"220":{"lat":"-30.03815057422600000","lng":"-51.21874713620700000"},"221":{"lat":"-30.03774757422600000","lng":"-51.21908213620700000"},"222":{"lat":"-30.03672957422600000","lng":"-51.21985313620700000"},"223":{"lat":"-30.03622257422600000","lng":"-51.22017213620700000"},"224":{"lat":"-30.03535357422600000","lng":"-51.22082413620700000"},"225":{"lat":"-30.03473757422600000","lng":"-51.22131613620700000"},"226":{"lat":"-30.03401857422600000","lng":"-51.22189913620700000"},"227":{"lat":"-30.03346957422600000","lng":"-51.22237913620700000"},"228":{"lat":"-30.03252457422600000","lng":"-51.22311713620700000"},"229":{"lat":"-30.03249957422600000","lng":"-51.22313613620700000"},"230":{"lat":"-30.03218957422600000","lng":"-51.22337813620700000"},"231":{"lat":"-30.03191857422600000","lng":"-51.22357813620700000"},"232":{"lat":"-30.03146857422600000","lng":"-51.22392513620700000"},"233":{"lat":"-30.03112757422600000","lng":"-51.22419013620700000"},"234":{"lat":"-30.03106957422600000","lng":"-51.22425713620700000"},"235":{"lat":"-30.03103157422600000","lng":"-51.22435313620700000"},"236":{"lat":"-30.03098657422600000","lng":"-51.22481213620700000"},"237":{"lat":"-30.03091357422600000","lng":"-51.22586813620700000"},"238":{"lat":"-30.03084257422600000","lng":"-51.22691413620700000"}}`
	got, err := jsonItinerarioDecode([]byte(req))
	if err != nil {
		t.Fatal("Falhei", err, got)
	}
}
