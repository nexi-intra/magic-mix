"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
import ColumnRead from "@/services/magic-mix/endpoints/column/read/webcomponent";
import { VsCodeEdittoolbar } from "@/app/magic/components/VsCodeEdittoolbar";

export default function TestColumnRead() {
return (
<div>
<VsCodeEdittoolbar
filePath={
  "app/magic/services/magic-mix/column/read/page.tsx"
}
/>
<ColumnRead />
</div>
);
}
    
