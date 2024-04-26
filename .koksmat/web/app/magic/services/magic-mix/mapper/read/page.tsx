"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
import MapperRead from "@/services/magic-mix/endpoints/mapper/read/webcomponent";
import { VsCodeEdittoolbar } from "@/app/magic/components/VsCodeEdittoolbar";

export default function TestMapperRead() {
return (
<div>
<VsCodeEdittoolbar
filePath={
  "app/magic/services/magic-mix/mapper/read/page.tsx"
}
/>
<MapperRead />
</div>
);
}
    
