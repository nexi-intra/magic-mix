"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
import IngressDelete from "@/services/magic-mix/endpoints/ingress/delete/webcomponent";
import { VsCodeEdittoolbar } from "@/app/magic/components/VsCodeEdittoolbar";

export default function TestIngressDelete() {
  return (
    <div>
      <VsCodeEdittoolbar
        filePath={"app/magic/services/magic-mix/ingress/delete/page.tsx"}
      />
      <IngressDelete />
    </div>
  );
}
