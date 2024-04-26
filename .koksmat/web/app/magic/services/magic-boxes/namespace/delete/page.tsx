"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
import NamespaceDelete from "@/services/magic-mix/endpoints/namespace/delete/webcomponent";
import { VsCodeEdittoolbar } from "@/app/magic/components/VsCodeEdittoolbar";

export default function TestNamespaceDelete() {
  return (
    <div>
      <VsCodeEdittoolbar
        filePath={"app/magic/services/magic-mix/namespace/delete/page.tsx"}
      />
      <NamespaceDelete />
    </div>
  );
}
