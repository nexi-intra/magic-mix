"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
import TransformerSearch from "@/services/magic-mix/endpoints/transformer/search/webcomponent";
import { VsCodeEdittoolbar } from "@/app/magic/components/VsCodeEdittoolbar";

export default function TestTransformerSearch() {
return (
<div>
<VsCodeEdittoolbar
filePath={
  "app/magic/services/magic-mix/transformer/search/page.tsx"
}
/>
<TransformerSearch />
</div>
);
}
    
