"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
import ServiceUpdate from "@/services/magic-mix/endpoints/service/update/webcomponent";
import { VsCodeEdittoolbar } from "@/app/magic/components/VsCodeEdittoolbar";

export default function TestServiceUpdate() {
  return (
    <div>
      <VsCodeEdittoolbar
        filePath={"app/magic/services/magic-mix/service/update/page.tsx"}
      />
      <ServiceUpdate />
    </div>
  );
}
