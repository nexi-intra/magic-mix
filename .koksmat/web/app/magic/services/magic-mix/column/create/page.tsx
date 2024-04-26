"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
import ColumnCreate from "@/services/magic-mix/endpoints/column/create/webcomponent";
import { VsCodeEdittoolbar } from "@/app/magic/components/VsCodeEdittoolbar";

export default function TestColumnCreate() {
return (
<div>
<VsCodeEdittoolbar
filePath={
  "app/magic/services/magic-mix/column/create/page.tsx"
}
/>
<ColumnCreate />
</div>
);
}
    
