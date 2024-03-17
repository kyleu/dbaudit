"use strict";(()=>{function l(e,n){let t;n?t=n.querySelectorAll(e):t=document.querySelectorAll(e);let o=[];return t.forEach(r=>{o.push(r)}),o}function k(e,n){let t=l(e,n);switch(t.length){case 0:return;case 1:return t[0];default:console.warn(`found [${t.length}] elements with selector [${e}], wanted zero or one`)}}function h(e,n){let t=k(e,n);if(!t)throw new Error(`no element found for selector [${e}]`);return t}function T(e,n,t="block"){return typeof e=="string"&&(e=h(e)),e.style.display=n?t:"none",e}function oe(e,...n){let t=document.createElement("li");t.innerText=e;for(let o of n){let r=document.createElement("pre");typeof o=="string"?r.innerText=o:r.innerText=JSON.stringify(o,null,2),t.appendChild(r)}return t}function O(e,...n){let t=k("#audit-log");t?t.appendChild(oe(e,...n)):console.log(`### Audit ###
`+e,...n)}function B(){for(let e of l(".menu-container .final"))e.scrollIntoView({block:"center"})}var x="mode-light",H="mode-dark";function q(){for(let e of l(".mode-input"))e.onclick=()=>{switch(e.value){case"":document.body.classList.remove(x),document.body.classList.remove(H);break;case"light":document.body.classList.add(x),document.body.classList.remove(H);break;case"dark":document.body.classList.remove(x),document.body.classList.add(H);break}}}function _(e){setTimeout(()=>{e.style.opacity="0",setTimeout(()=>e.remove(),500)},5e3)}function R(e,n,t){let o=document.getElementById("flash-container");o===null&&(o=document.createElement("div"),o.id="flash-container",document.body.insertAdjacentElement("afterbegin",o));let r=document.createElement("div");r.className="flash";let i=document.createElement("input");i.type="radio",i.style.display="none",i.id="hide-flash-"+e,r.appendChild(i);let s=document.createElement("label");s.htmlFor="hide-flash-"+e;let c=document.createElement("span");c.innerHTML="\xD7",s.appendChild(c),r.appendChild(s);let a=document.createElement("div");a.className="content flash-"+n,a.innerText=t,r.appendChild(a),o.appendChild(r),_(r)}function $(){let e=document.getElementById("flash-container");if(e===null)return R;let n=e.querySelectorAll(".flash");if(n.length>0)for(let t of n)_(t);return R}function j(){for(let e of l(".link-confirm"))e.onclick=()=>{let n=e.dataset.message;return n&&n.length===0&&(n="Are you sure?"),confirm(n)}}function G(e){let n=e.dataset.timestamp;if(n===""){console.log("invalid timestamp ["+n+"]",e);return}let t=new Date(n),o=e.dataset.format;o===""&&(o="yyyy-MM-dd hh:mm:ss");let r=document.createElement("span");r.title=t.toISOString(),r.innerText=t.toLocaleString(),e.replaceChildren(r)}function w(e){let t=(e.dataset.timestamp||"").replace(/-/gu,"/").replace(/[TZ]/gu," ")+" UTC",o=new Date(t),r=(new Date().getTime()-o.getTime())/1e3,i=Math.floor(r/86400),s=o.getFullYear(),c=o.getMonth()+1,a=o.getDate();if(isNaN(i)||i<0||i>=31)return s.toString()+"-"+(c<10?"0"+c.toString():c.toString())+"-"+(a<10?"0"+a.toString():a.toString());let m,u;return i===0?r<5?(u=1,m="just now"):r<60?(u=1,m=Math.floor(r)+" seconds ago"):r<120?(u=10,m="1 minute ago"):r<3600?(u=30,m=Math.floor(r/60)+" minutes ago"):r<7200?(u=60,m="1 hour ago"):(u=60,m=Math.floor(r/3600)+" hours ago"):i===1?(u=600,m="yesterday"):i<7?(u=600,m=i+" days ago"):(u=6e3,m=Math.ceil(i/7)+" weeks ago"),e&&(e.innerText=m,setTimeout(()=>w(e),u*1e3)),m}function W(){return l(".timestamp").forEach(e=>{G(e)}),l(".reltime").forEach(e=>{w(e)}),[G,w]}function re(e,n){let t=0;return(...o)=>{t!==0&&window.clearTimeout(t),t=window.setTimeout(()=>{e(null,...o)},n)}}function ie(e,n,t,o,r){var D;if(!e)return;let i=e.id+"-list",s=document.createElement("datalist"),c=document.createElement("option");c.value="",c.innerText="Loading...",s.appendChild(c),s.id=i,(D=e.parentElement)==null||D.prepend(s),e.setAttribute("autocomplete","off"),e.setAttribute("list",i);let a={},m="";function u(d){let f=n;return f.includes("?")?f+"&t=json&"+t+"="+encodeURIComponent(d):f+"?t=json&"+t+"="+encodeURIComponent(d)}function N(d){let f=a[d];!f||!f.frag||(m=d,s.replaceChildren(f.frag.cloneNode(!0)))}function te(){let d=e.value;if(d.length===0)return;let f=u(d),b=!d||!m;if(!b){let p=a[m];p&&(b=!p.data.find(E=>d===r(E)))}if(b){if(a[d]&&a[d].url===f){N(d);return}fetch(f,{credentials:"include"}).then(p=>p.json()).then(p=>{if(!p)return;let E=Array.isArray(p)?p:[p],A=document.createDocumentFragment(),F=10;for(let v=0;v<E.length&&F>0;v++){let U=r(E[v]),ne=o(E[v]);if(U){let M=document.createElement("option");M.value=U,M.innerText=ne,A.appendChild(M),F--}}a[d]={url:f,data:E,frag:A,complete:!1},N(d)})}}e.oninput=re(te,250),console.log("managing ["+e.id+"] autocomplete")}function P(){return ie}function Y(){document.addEventListener("keydown",e=>{e.key==="Escape"&&document.location.hash.startsWith("#modal-")&&(document.location.hash="")})}function C(e,n,t){return n||(n=18),t==null&&(t="icon"),`<svg class="${t}" style="width: ${n}px; height: ${n+"px"};"><use xlink:href="#svg-${e}"></use></svg>`}function se(e,n){return e.parentElement!==n.parentElement?null:e===n?0:e.compareDocumentPosition(n)&Node.DOCUMENT_POSITION_FOLLOWING?-1:1}var y;function I(e){let n=[],t=l(".item .value",e);for(let r of t)n.push(r.innerText);let o=h("input.result",e);o.value=n.join(", ")}function le(e){var t;let n=(t=e.parentElement)==null?void 0:t.parentElement;e.remove(),n&&I(n)}function K(e){let n=h(".value",e),t=h(".editor",e);t.value=n.innerText;let o=()=>{var i;if(t.value===""){e.remove();return}n.innerText=t.value,T(n,!0),T(t,!1);let r=(i=e.parentElement)==null?void 0:i.parentElement;r&&I(r)};t.onblur=o,t.onkeydown=r=>{if(r.code==="Enter")return r.preventDefault(),o(),!1},T(n,!1),T(t,!0),t.focus()}function Q(e,n){let t=document.createElement("div");t.className="item",t.draggable=!0,t.ondragstart=s=>{var c;(c=s.dataTransfer)==null||c.setDragImage(document.createElement("div"),0,0),t.classList.add("dragging"),y=t},t.ondragover=()=>{var a;let s=se(t,y);if(!s)return;let c=s===-1?t:t.nextSibling;(a=y.parentElement)==null||a.insertBefore(y,c),I(n)},t.ondrop=s=>{s.preventDefault()},t.ondragend=s=>{t.classList.remove("dragging"),s.preventDefault()};let o=document.createElement("div");o.innerText=e,o.className="value",o.onclick=()=>{K(t)},t.appendChild(o);let r=document.createElement("input");r.className="editor",t.appendChild(r);let i=document.createElement("div");return i.innerHTML=C("times",13),i.className="close",i.onclick=()=>{le(t)},t.appendChild(i),t}function ae(e,n){let t=Q("",n);e.appendChild(t),K(t)}function J(e){var i;let n=h("input.result",e),t=h(".tags",e),o=n.value.split(",").map(s=>s.trim()).filter(s=>s!=="");T(n,!1),t.innerHTML="";for(let s of o)t.appendChild(Q(s,e));(i=k(".add-item",e))==null||i.remove();let r=document.createElement("div");r.className="add-item",r.innerHTML=C("plus",22),r.onclick=()=>{ae(t,e)},e.insertBefore(r,h(".clear",e))}function V(){for(let e of l(".tag-editor"))J(e);return J}var Z="--selected";function ce(e){var t,o;let n=(o=(t=e.parentElement)==null?void 0:t.parentElement)==null?void 0:o.querySelector("input");if(!n)throw new Error("no associated input found");n.value="\u2205"}function S(e){e.onreset=()=>S(e);let n={},t={};for(let o of e.elements){let r=o;if(r.name.length>0)if(r.name.endsWith(Z))t[r.name]=r;else{(r.type!=="radio"||r.checked)&&(n[r.name]=r.value);let i=()=>{let s=t[r.name+Z];s&&(s.checked=n[r.name]!==r.value)};r.onchange=i,r.onkeyup=i}}}function X(){for(let e of l("form.editor"))S(e);return[ce,S]}var me=[];function z(e,n,t){let o=e.querySelectorAll(n);if(o.length===0)throw new Error("empty query selector ["+n+"]");o.forEach(r=>t(r))}function L(e,n,t){z(e,n,o=>{o.style.backgroundColor=t})}function g(e,n,t){z(e,n,o=>{o.style.color=t})}function de(e,n,t){let o=document.querySelector("#mockup-"+e);if(!o){console.error("can't find mockup for mode ["+e+"]");return}switch(n){case"color-foreground":g(o,".mock-main",t);break;case"color-background":L(o,".mock-main",t);break;case"color-foreground-muted":g(o,".mock-main .mock-muted",t);break;case"color-background-muted":L(o,".mock-main .mock-muted",t);break;case"color-link-foreground":g(o,".mock-main .mock-link",t);break;case"color-link-visited-foreground":g(o,".mock-main .mock-link-visited",t);break;case"color-nav-foreground":g(o,".mock-nav",t),g(o,".mock-nav .mock-link",t);break;case"color-nav-background":L(o,".mock-nav",t);break;case"color-menu-foreground":g(o,".mock-menu",t),g(o,".mock-menu .mock-link",t);break;case"color-menu-background":L(o,".mock-menu",t);break;case"color-menu-selected-foreground":g(o,".mock-menu .mock-link-selected",t);break;case"color-menu-selected-background":L(o,".mock-menu .mock-link-selected",t);break;default:console.error("invalid key ["+n+"]")}}function ee(){for(let e of l(".color-var")){let n=e.dataset.var,t=e.dataset.mode;me.push(n),!(!n||n.length===0)&&(e.oninput=()=>{de(t,n,e.value)})}}function ue(){let[e,n]=X(),[t,o]=W();window.dbaudit={wireTime:t,relativeTime:o,autocomplete:P(),setSiblingToNull:e,initForm:n,flash:$(),tags:V()},B(),q(),j(),Y(),ee(),window.audit=O}document.addEventListener("DOMContentLoaded",ue);})();
//# sourceMappingURL=client.js.map
