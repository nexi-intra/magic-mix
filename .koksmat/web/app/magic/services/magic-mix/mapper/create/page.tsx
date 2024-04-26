"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
import MapperCreate from "@/services/magic-mix/endpoints/mapper/create/webcomponent";
import { VsCodeEdittoolbar } from "@/app/magic/components/VsCodeEdittoolbar";

export default function TestMapperCreate() {
return (
<div>
<VsCodeEdittoolbar
filePath={
  "app/magic/services/magic-mix/mapper/create/page.tsx"
}
/>
<MapperCreate />
</div>
);
}
    
