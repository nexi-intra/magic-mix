"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
import ProcessLogCreate from "@/services/magic-mix/endpoints/processlog/create/webcomponent";
import { VsCodeEdittoolbar } from "@/app/magic/components/VsCodeEdittoolbar";

export default function TestProcessLogCreate() {
return (
<div>
<VsCodeEdittoolbar
filePath={
  "app/magic/services/magic-mix/processlog/create/page.tsx"
}
/>
<ProcessLogCreate />
</div>
);
}
    
