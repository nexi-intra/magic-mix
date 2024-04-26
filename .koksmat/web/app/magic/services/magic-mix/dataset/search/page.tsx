"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
import DatasetSearch from "@/services/magic-mix/endpoints/dataset/search/webcomponent";
import { VsCodeEdittoolbar } from "@/app/magic/components/VsCodeEdittoolbar";

export default function TestDatasetSearch() {
return (
<div>
<VsCodeEdittoolbar
filePath={
  "app/magic/services/magic-mix/dataset/search/page.tsx"
}
/>
<DatasetSearch />
</div>
);
}
    
