"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
import TransformationRead from "@/services/magic-mix/endpoints/transformation/read/webcomponent";
import { VsCodeEdittoolbar } from "@/app/magic/components/VsCodeEdittoolbar";

export default function TestTransformationRead() {
return (
<div>
<VsCodeEdittoolbar
filePath={
  "app/magic/services/magic-mix/transformation/read/page.tsx"
}
/>
<TransformationRead />
</div>
);
}
    
