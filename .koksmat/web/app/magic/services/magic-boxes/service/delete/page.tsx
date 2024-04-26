"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
import ServiceDelete from "@/services/magic-mix/endpoints/service/delete/webcomponent";
import { VsCodeEdittoolbar } from "@/app/magic/components/VsCodeEdittoolbar";

export default function TestServiceDelete() {
  return (
    <div>
      <VsCodeEdittoolbar
        filePath={"app/magic/services/magic-mix/service/delete/page.tsx"}
      />
      <ServiceDelete />
    </div>
  );
}
