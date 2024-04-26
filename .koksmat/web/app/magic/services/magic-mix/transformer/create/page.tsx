"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
import TransformerCreate from "@/services/magic-mix/endpoints/transformer/create/webcomponent";
import { VsCodeEdittoolbar } from "@/app/magic/components/VsCodeEdittoolbar";

export default function TestTransformerCreate() {
return (
<div>
<VsCodeEdittoolbar
filePath={
  "app/magic/services/magic-mix/transformer/create/page.tsx"
}
/>
<TransformerCreate />
</div>
);
}
    
