"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
import ProcessLogSearch from "@/services/magic-mix/endpoints/processlog/search/webcomponent";
import { VsCodeEdittoolbar } from "@/app/magic/components/VsCodeEdittoolbar";

export default function TestProcessLogSearch() {
return (
<div>
<VsCodeEdittoolbar
filePath={
  "app/magic/services/magic-mix/processlog/search/page.tsx"
}
/>
<ProcessLogSearch />
</div>
);
}
    
