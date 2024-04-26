"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
import ConnectionSearch from "@/services/magic-mix/endpoints/connection/search/webcomponent";
import { VsCodeEdittoolbar } from "@/app/magic/components/VsCodeEdittoolbar";

export default function TestConnectionSearch() {
return (
<div>
<VsCodeEdittoolbar
filePath={
  "app/magic/services/magic-mix/connection/search/page.tsx"
}
/>
<ConnectionSearch />
</div>
);
}
    
