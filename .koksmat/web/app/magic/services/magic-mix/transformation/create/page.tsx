"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
import TransformationCreate from "@/services/magic-mix/endpoints/transformation/create/webcomponent";
import { VsCodeEdittoolbar } from "@/app/magic/components/VsCodeEdittoolbar";

export default function TestTransformationCreate() {
return (
<div>
<VsCodeEdittoolbar
filePath={
  "app/magic/services/magic-mix/transformation/create/page.tsx"
}
/>
<TransformationCreate />
</div>
);
}
    
