"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
import EnvironmentDelete from "@/services/magic-mix/endpoints/environment/delete/webcomponent";
import { VsCodeEdittoolbar } from "@/app/magic/components/VsCodeEdittoolbar";

export default function TestEnvironmentDelete() {
  return (
    <div>
      <VsCodeEdittoolbar
        filePath={"app/magic/services/magic-mix/environment/delete/page.tsx"}
      />
      <EnvironmentDelete />
    </div>
  );
}
