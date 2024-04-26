"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
import ProcessLogRead from "@/services/magic-mix/endpoints/processlog/read/webcomponent";
import { VsCodeEdittoolbar } from "@/app/magic/components/VsCodeEdittoolbar";

export default function TestProcessLogRead() {
return (
<div>
<VsCodeEdittoolbar
filePath={
  "app/magic/services/magic-mix/processlog/read/page.tsx"
}
/>
<ProcessLogRead />
</div>
);
}
    
