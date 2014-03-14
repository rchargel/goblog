document.createElement('ng-include');document.createElement('ng-pluralize');document.createElement('ng-view');document.createElement('ng:include');document.createElement('ng:pluralize');document.createElement('ng:view');/*
 HTML5 Shiv v3.6.2pre | @afarkas @jdalton @jon_neal @rem | MIT/GPL2 Licensed
*/
(function(r,f){function s(){var a=g.elements;return"string"==typeof a?a.split(" "):a}function j(a){var b=t[a[u]];b||(b={},k++,a[u]=k,t[k]=b);return b}function v(a,b,c){b||(b=f);if(h)return b.createElement(a);c||(c=j(b));b=c.cache[a]?c.cache[a].cloneNode():x.test(a)?(c.cache[a]=c.createElem(a)).cloneNode():c.createElem(a);return b.canHaveChildren&&!y.test(a)?c.frag.appendChild(b):b}function w(a){a||(a=f);var b=j(a);if(g.shivCSS&&!l&&!b.hasCSS){var c,d=a;c=d.createElement("p");d=d.getElementsByTagName("head")[0]||
d.documentElement;c.innerHTML="x<style>article,aside,figcaption,figure,footer,header,hgroup,nav,section{display:block}mark{background:#FF0;color:#000}</style>";c=d.insertBefore(c.lastChild,d.firstChild);b.hasCSS=!!c}if(!h){var e=a;b.cache||(b.cache={},b.createElem=e.createElement,b.createFrag=e.createDocumentFragment,b.frag=b.createFrag());e.createElement=function(a){return!g.shivMethods?b.createElem(a):v(a,e,b)};e.createDocumentFragment=Function("h,f","return function(){var n=f.cloneNode(),c=n.createElement;h.shivMethods&&("+
s().join().replace(/\w+/g,function(a){b.createElem(a);b.frag.createElement(a);return'c("'+a+'")'})+");return n}")(g,b.frag)}return a}var m=r.html5||{},y=/^<|^(?:button|map|select|textarea|object|iframe|option|optgroup)$/i,x=/^(?:a|b|code|div|fieldset|h1|h2|h3|h4|h5|h6|i|label|li|ol|p|q|span|strong|style|table|tbody|td|th|tr|ul)$/i,l,u="_html5shiv",k=0,t={},h;try{var n=f.createElement("a");n.innerHTML="<xyz></xyz>";l="hidden"in n;var p;if(!(p=1==n.childNodes.length)){f.createElement("a");var q=f.createDocumentFragment();
p="undefined"==typeof q.cloneNode||"undefined"==typeof q.createDocumentFragment||"undefined"==typeof q.createElement}h=p}catch(z){h=l=!0}var g={elements:m.elements||"abbr article aside audio bdi canvas data datalist details figcaption figure footer header hgroup main mark meter nav output progress section summary time video",version:"3.6.2pre",shivCSS:!1!==m.shivCSS,supportsUnknownElements:h,shivMethods:!1!==m.shivMethods,type:"default",shivDocument:w,createElement:v,createDocumentFragment:function(a,
b){a||(a=f);if(h)return a.createDocumentFragment();b=b||j(a);for(var c=b.frag.cloneNode(),d=0,e=s(),g=e.length;d<g;d++)c.createElement(e[d]);return c}};r.html5=g;w(f)})(this,document);
(function(){var getClass={}.toString,isProperty,forEach,undef;var isLoader=typeof define==="function"&&define.amd,JSON3=!isLoader&&typeof exports=="object"&&exports;if(JSON3||isLoader)if(typeof JSON=="object"&&JSON)if(isLoader)JSON3=JSON;else{JSON3.stringify=JSON.stringify;JSON3.parse=JSON.parse}else{if(isLoader)JSON3=this.JSON={}}else JSON3=this.JSON||(this.JSON={});var Escapes,toPaddedString,quote,serialize;var fromCharCode,Unescapes,abort,lex,get,walk,update,Index,Source;var isExtended=new Date(-0xc782b5b800cec),
floor,Months,getDay;try{isExtended=isExtended.getUTCFullYear()==-109252&&isExtended.getUTCMonth()===0&&isExtended.getUTCDate()==1&&isExtended.getUTCHours()==10&&isExtended.getUTCMinutes()==37&&isExtended.getUTCSeconds()==6&&isExtended.getUTCMilliseconds()==708}catch(exception){}function has(name){var stringifySupported,parseSupported,value,serialized='{"A":[1,true,false,null,"\\u0000\\b\\n\\f\\r\\t"]}',all=name=="json";if(all||name=="json-stringify"||name=="json-parse"){if(name=="json-stringify"||
all){if(stringifySupported=typeof JSON3.stringify=="function"&&isExtended){(value=function(){return 1}).toJSON=value;try{stringifySupported=JSON3.stringify(0)==="0"&&JSON3.stringify(new Number)==="0"&&JSON3.stringify(new String)=='""'&&JSON3.stringify(getClass)===undef&&JSON3.stringify(undef)===undef&&JSON3.stringify()===undef&&JSON3.stringify(value)==="1"&&JSON3.stringify([value])=="[1]"&&JSON3.stringify([undef])=="[null]"&&JSON3.stringify(null)=="null"&&JSON3.stringify([undef,getClass,null])=="[null,null,null]"&&
JSON3.stringify({"A":[value,true,false,null,"\x00\b\n\f\r\t"]})==serialized&&JSON3.stringify(null,value)==="1"&&JSON3.stringify([1,2],null,1)=="[\n 1,\n 2\n]"&&JSON3.stringify(new Date(-864E13))=='"-271821-04-20T00:00:00.000Z"'&&JSON3.stringify(new Date(864E13))=='"+275760-09-13T00:00:00.000Z"'&&JSON3.stringify(new Date(-621987552E5))=='"-000001-01-01T00:00:00.000Z"'&&JSON3.stringify(new Date(-1))=='"1969-12-31T23:59:59.999Z"'}catch(exception){stringifySupported=false}}if(!all)return stringifySupported}if(name==
"json-parse"||all){if(typeof JSON3.parse=="function")try{if(JSON3.parse("0")===0&&!JSON3.parse(false)){value=JSON3.parse(serialized);if(parseSupported=value.A.length==5&&value.A[0]==1){try{parseSupported=!JSON3.parse('"\t"')}catch(exception){}if(parseSupported)try{parseSupported=JSON3.parse("01")!=1}catch(exception){}}}}catch(exception){parseSupported=false}if(!all)return parseSupported}return stringifySupported&&parseSupported}}if(!has("json")){if(!isExtended){floor=Math.floor;Months=[0,31,59,90,
120,151,181,212,243,273,304,334];getDay=function(year,month){return Months[month]+365*(year-1970)+floor((year-1969+(month=+(month>1)))/4)-floor((year-1901+month)/100)+floor((year-1601+month)/400)}}if(!(isProperty={}.hasOwnProperty))isProperty=function(property){var members={},constructor;if((members.__proto__=null,members.__proto__={"toString":1},members).toString!=getClass)isProperty=function(property){var original=this.__proto__,result=property in(this.__proto__=null,this);this.__proto__=original;
return result};else{constructor=members.constructor;isProperty=function(property){var parent=(this.constructor||constructor).prototype;return property in this&&!(property in parent&&this[property]===parent[property])}}members=null;return isProperty.call(this,property)};forEach=function(object,callback){var size=0,Properties,members,property,forEach;(Properties=function(){this.valueOf=0}).prototype.valueOf=0;members=new Properties;for(property in members)if(isProperty.call(members,property))size++;
Properties=members=null;if(!size){members=["valueOf","toString","toLocaleString","propertyIsEnumerable","isPrototypeOf","hasOwnProperty","constructor"];forEach=function(object,callback){var isFunction=getClass.call(object)=="[object Function]",property,length;for(property in object)if(!(isFunction&&property=="prototype")&&isProperty.call(object,property))callback(property);for(length=members.length;property=members[--length];isProperty.call(object,property)&&callback(property));}}else if(size==2)forEach=
function(object,callback){var members={},isFunction=getClass.call(object)=="[object Function]",property;for(property in object)if(!(isFunction&&property=="prototype")&&!isProperty.call(members,property)&&(members[property]=1)&&isProperty.call(object,property))callback(property)};else forEach=function(object,callback){var isFunction=getClass.call(object)=="[object Function]",property,isConstructor;for(property in object)if(!(isFunction&&property=="prototype")&&isProperty.call(object,property)&&!(isConstructor=
property==="constructor"))callback(property);if(isConstructor||isProperty.call(object,property="constructor"))callback(property)};return forEach(object,callback)};if(!has("json-stringify")){Escapes={"\\":"\\\\",'"':'\\"',"\b":"\\b","\f":"\\f","\n":"\\n","\r":"\\r","\t":"\\t"};toPaddedString=function(width,value){return("000000"+(value||0)).slice(-width)};quote=function(value){var result='"',index=0,symbol;for(;symbol=value.charAt(index);index++)result+='\\"\b\f\n\r\t'.indexOf(symbol)>-1?Escapes[symbol]:
Escapes[symbol]=symbol<" "?"\\u00"+toPaddedString(2,symbol.charCodeAt(0).toString(16)):symbol;return result+'"'};serialize=function(property,object,callback,properties,whitespace,indentation,stack){var value=object[property],className,year,month,date,time,hours,minutes,seconds,milliseconds,results,element,index,length,prefix,any,result;if(typeof value=="object"&&value){className=getClass.call(value);if(className=="[object Date]"&&!isProperty.call(value,"toJSON"))if(value>-1/0&&value<1/0){if(getDay){date=
floor(value/864E5);for(year=floor(date/365.2425)+1970-1;getDay(year+1,0)<=date;year++);for(month=floor((date-getDay(year,0))/30.42);getDay(year,month+1)<=date;month++);date=1+date-getDay(year,month);time=(value%864E5+864E5)%864E5;hours=floor(time/36E5)%24;minutes=floor(time/6E4)%60;seconds=floor(time/1E3)%60;milliseconds=time%1E3}else{year=value.getUTCFullYear();month=value.getUTCMonth();date=value.getUTCDate();hours=value.getUTCHours();minutes=value.getUTCMinutes();seconds=value.getUTCSeconds();
milliseconds=value.getUTCMilliseconds()}value=(year<=0||year>=1E4?(year<0?"-":"+")+toPaddedString(6,year<0?-year:year):toPaddedString(4,year))+"-"+toPaddedString(2,month+1)+"-"+toPaddedString(2,date)+"T"+toPaddedString(2,hours)+":"+toPaddedString(2,minutes)+":"+toPaddedString(2,seconds)+"."+toPaddedString(3,milliseconds)+"Z"}else value=null;else if(typeof value.toJSON=="function"&&(className!="[object Number]"&&className!="[object String]"&&className!="[object Array]"||isProperty.call(value,"toJSON")))value=
value.toJSON(property)}if(callback)value=callback.call(object,property,value);if(value===null)return"null";className=getClass.call(value);if(className=="[object Boolean]")return""+value;else if(className=="[object Number]")return value>-1/0&&value<1/0?""+value:"null";else if(className=="[object String]")return quote(value);if(typeof value=="object"){for(length=stack.length;length--;)if(stack[length]===value)throw TypeError();stack.push(value);results=[];prefix=indentation;indentation+=whitespace;
if(className=="[object Array]"){for(index=0,length=value.length;index<length;any||(any=true),index++){element=serialize(index,value,callback,properties,whitespace,indentation,stack);results.push(element===undef?"null":element)}result=any?whitespace?"[\n"+indentation+results.join(",\n"+indentation)+"\n"+prefix+"]":"["+results.join(",")+"]":"[]"}else{forEach(properties||value,function(property){var element=serialize(property,value,callback,properties,whitespace,indentation,stack);if(element!==undef)results.push(quote(property)+
":"+(whitespace?" ":"")+element);any||(any=true)});result=any?whitespace?"{\n"+indentation+results.join(",\n"+indentation)+"\n"+prefix+"}":"{"+results.join(",")+"}":"{}"}stack.pop();return result}};JSON3.stringify=function(source,filter,width){var whitespace,callback,properties,index,length,value;if(typeof filter=="function"||typeof filter=="object"&&filter)if(getClass.call(filter)=="[object Function]")callback=filter;else if(getClass.call(filter)=="[object Array]"){properties={};for(index=0,length=
filter.length;index<length;value=filter[index++],(getClass.call(value)=="[object String]"||getClass.call(value)=="[object Number]")&&(properties[value]=1));}if(width)if(getClass.call(width)=="[object Number]"){if((width-=width%1)>0)for(whitespace="",width>10&&(width=10);whitespace.length<width;whitespace+=" ");}else if(getClass.call(width)=="[object String]")whitespace=width.length<=10?width:width.slice(0,10);return serialize("",(value={},value[""]=source,value),callback,properties,whitespace,"",
[])}}if(!has("json-parse")){fromCharCode=String.fromCharCode;Unescapes={"\\":"\\",'"':'"',"/":"/","b":"\b","t":"\t","n":"\n","f":"\f","r":"\r"};abort=function(){Index=Source=null;throw SyntaxError();};lex=function(){var source=Source,length=source.length,symbol,value,begin,position,sign;while(Index<length){symbol=source.charAt(Index);if("\t\r\n ".indexOf(symbol)>-1)Index++;else if("{}[]:,".indexOf(symbol)>-1){Index++;return symbol}else if(symbol=='"'){for(value="@",Index++;Index<length;){symbol=source.charAt(Index);
if(symbol<" ")abort();else if(symbol=="\\"){symbol=source.charAt(++Index);if('\\"/btnfr'.indexOf(symbol)>-1){value+=Unescapes[symbol];Index++}else if(symbol=="u"){begin=++Index;for(position=Index+4;Index<position;Index++){symbol=source.charAt(Index);if(!(symbol>="0"&&symbol<="9"||symbol>="a"&&symbol<="f"||symbol>="A"&&symbol<="F"))abort()}value+=fromCharCode("0x"+source.slice(begin,Index))}else abort()}else{if(symbol=='"')break;value+=symbol;Index++}}if(source.charAt(Index)=='"'){Index++;return value}abort()}else{begin=
Index;if(symbol=="-"){sign=true;symbol=source.charAt(++Index)}if(symbol>="0"&&symbol<="9"){if(symbol=="0"&&(symbol=source.charAt(Index+1),symbol>="0"&&symbol<="9"))abort();sign=false;for(;Index<length&&(symbol=source.charAt(Index),symbol>="0"&&symbol<="9");Index++);if(source.charAt(Index)=="."){position=++Index;for(;position<length&&(symbol=source.charAt(position),symbol>="0"&&symbol<="9");position++);if(position==Index)abort();Index=position}symbol=source.charAt(Index);if(symbol=="e"||symbol=="E"){symbol=
source.charAt(++Index);if(symbol=="+"||symbol=="-")Index++;for(position=Index;position<length&&(symbol=source.charAt(position),symbol>="0"&&symbol<="9");position++);if(position==Index)abort();Index=position}return+source.slice(begin,Index)}if(sign)abort();if(source.slice(Index,Index+4)=="true"){Index+=4;return true}else if(source.slice(Index,Index+5)=="false"){Index+=5;return false}else if(source.slice(Index,Index+4)=="null"){Index+=4;return null}abort()}}return"$"};get=function(value){var results,
any,key;if(value=="$")abort();if(typeof value=="string"){if(value.charAt(0)=="@")return value.slice(1);if(value=="["){results=[];for(;;any||(any=true)){value=lex();if(value=="]")break;if(any)if(value==","){value=lex();if(value=="]")abort()}else abort();if(value==",")abort();results.push(get(value))}return results}else if(value=="{"){results={};for(;;any||(any=true)){value=lex();if(value=="}")break;if(any)if(value==","){value=lex();if(value=="}")abort()}else abort();if(value==","||typeof value!="string"||
value.charAt(0)!="@"||lex()!=":")abort();results[value.slice(1)]=get(lex())}return results}abort()}return value};update=function(source,property,callback){var element=walk(source,property,callback);if(element===undef)delete source[property];else source[property]=element};walk=function(source,property,callback){var value=source[property],length;if(typeof value=="object"&&value)if(getClass.call(value)=="[object Array]")for(length=value.length;length--;)update(value,length,callback);else forEach(value,
function(property){update(value,property,callback)});return callback.call(source,property,value)};JSON3.parse=function(source,callback){var result,value;Index=0;Source=source;result=get(lex());if(lex()!="$")abort();Index=Source=null;return callback&&getClass.call(callback)=="[object Function]"?walk((value={},value[""]=result,value),"",callback):result}}}if(isLoader)define(function(){return JSON3})}).call(this);
/*! matchMedia() polyfill - Test a CSS media type/query in JS. Authors & copyright (c) 2012: Scott Jehl, Paul Irish, Nicholas Zakas. Dual MIT/BSD license */
/*! NOTE: If you're already including a window.matchMedia polyfill via Modernizr or otherwise, you don't need this part */
window.matchMedia=window.matchMedia||function(a){"use strict";var c,d=a.documentElement,e=d.firstElementChild||d.firstChild,f=a.createElement("body"),g=a.createElement("div");return g.id="mq-test-1",g.style.cssText="position:absolute;top:-100em",f.style.background="none",f.appendChild(g),function(a){return g.innerHTML='&shy;<style media="'+a+'"> #mq-test-1 { width: 42px; }</style>',d.insertBefore(f,e),c=42===g.offsetWidth,d.removeChild(f),{matches:c,media:a}}}(document);

/*! Respond.js v1.1.0: min/max-width media query polyfill. (c) Scott Jehl. MIT/GPLv2 Lic. j.mp/respondjs  */
(function(a){"use strict";function x(){u(!0)}var b={};if(a.respond=b,b.update=function(){},b.mediaQueriesSupported=a.matchMedia&&a.matchMedia("only all").matches,!b.mediaQueriesSupported){var q,r,t,c=a.document,d=c.documentElement,e=[],f=[],g=[],h={},i=30,j=c.getElementsByTagName("head")[0]||d,k=c.getElementsByTagName("base")[0],l=j.getElementsByTagName("link"),m=[],n=function(){for(var b=0;l.length>b;b++){var c=l[b],d=c.href,e=c.media,f=c.rel&&"stylesheet"===c.rel.toLowerCase();d&&f&&!h[d]&&(c.styleSheet&&c.styleSheet.rawCssText?(p(c.styleSheet.rawCssText,d,e),h[d]=!0):(!/^([a-zA-Z:]*\/\/)/.test(d)&&!k||d.replace(RegExp.$1,"").split("/")[0]===a.location.host)&&m.push({href:d,media:e}))}o()},o=function(){if(m.length){var b=m.shift();v(b.href,function(c){p(c,b.href,b.media),h[b.href]=!0,a.setTimeout(function(){o()},0)})}},p=function(a,b,c){var d=a.match(/@media[^\{]+\{([^\{\}]*\{[^\}\{]*\})+/gi),g=d&&d.length||0;b=b.substring(0,b.lastIndexOf("/"));var h=function(a){return a.replace(/(url\()['"]?([^\/\)'"][^:\)'"]+)['"]?(\))/g,"$1"+b+"$2$3")},i=!g&&c;b.length&&(b+="/"),i&&(g=1);for(var j=0;g>j;j++){var k,l,m,n;i?(k=c,f.push(h(a))):(k=d[j].match(/@media *([^\{]+)\{([\S\s]+?)$/)&&RegExp.$1,f.push(RegExp.$2&&h(RegExp.$2))),m=k.split(","),n=m.length;for(var o=0;n>o;o++)l=m[o],e.push({media:l.split("(")[0].match(/(only\s+)?([a-zA-Z]+)\s?/)&&RegExp.$2||"all",rules:f.length-1,hasquery:l.indexOf("(")>-1,minw:l.match(/\(\s*min\-width\s*:\s*(\s*[0-9\.]+)(px|em)\s*\)/)&&parseFloat(RegExp.$1)+(RegExp.$2||""),maxw:l.match(/\(\s*max\-width\s*:\s*(\s*[0-9\.]+)(px|em)\s*\)/)&&parseFloat(RegExp.$1)+(RegExp.$2||"")})}u()},s=function(){var a,b=c.createElement("div"),e=c.body,f=!1;return b.style.cssText="position:absolute;font-size:1em;width:1em",e||(e=f=c.createElement("body"),e.style.background="none"),e.appendChild(b),d.insertBefore(e,d.firstChild),a=b.offsetWidth,f?d.removeChild(e):e.removeChild(b),a=t=parseFloat(a)},u=function(b){var h="clientWidth",k=d[h],m="CSS1Compat"===c.compatMode&&k||c.body[h]||k,n={},o=l[l.length-1],p=(new Date).getTime();if(b&&q&&i>p-q)return a.clearTimeout(r),r=a.setTimeout(u,i),void 0;q=p;for(var v in e)if(e.hasOwnProperty(v)){var w=e[v],x=w.minw,y=w.maxw,z=null===x,A=null===y,B="em";x&&(x=parseFloat(x)*(x.indexOf(B)>-1?t||s():1)),y&&(y=parseFloat(y)*(y.indexOf(B)>-1?t||s():1)),w.hasquery&&(z&&A||!(z||m>=x)||!(A||y>=m))||(n[w.media]||(n[w.media]=[]),n[w.media].push(f[w.rules]))}for(var C in g)g.hasOwnProperty(C)&&g[C]&&g[C].parentNode===j&&j.removeChild(g[C]);for(var D in n)if(n.hasOwnProperty(D)){var E=c.createElement("style"),F=n[D].join("\n");E.type="text/css",E.media=D,j.insertBefore(E,o.nextSibling),E.styleSheet?E.styleSheet.cssText=F:E.appendChild(c.createTextNode(F)),g.push(E)}},v=function(a,b){var c=w();c&&(c.open("GET",a,!0),c.onreadystatechange=function(){4!==c.readyState||200!==c.status&&304!==c.status||b(c.responseText)},4!==c.readyState&&c.send(null))},w=function(){var b=!1;try{b=new a.XMLHttpRequest}catch(c){b=new a.ActiveXObject("Microsoft.XMLHTTP")}return function(){return b}}();n(),b.update=n,a.addEventListener?a.addEventListener("resize",x,!1):a.attachEvent&&a.attachEvent("onresize",x)}})(this);