import path from "path";
import { cwd } from "process";
import fs from "fs";
import { pagemap } from "./magic-mix";
import { TestServicesNavigator } from "../components/testservicenavigator";
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";
import { MenuIcon } from "lucide-react";
import { OpenInNewWindowIcon } from "@radix-ui/react-icons";
import { VsCodeEdittoolbar } from "../components/VsCodeEdittoolbar";

export default function Layout(props: { children: any }) {
  const { children } = props;
  const folderPath = path.join(cwd(), "app", "test", "services");
  if (process.env.NODE_ENV === "production")
    return (
      <div className="text-orange-400">Only works in development mode</div>
    );

  return (
    <div>
      <div className="p-10">
        <h1 className="text-2xl">Services Test pages</h1>
        {/* <div>{folderPath}</div> */}
        <div className="flex">
          <div className="hidden lg:block lg:w-1/4">
            <TestServicesNavigator appMap={pagemap} appName="magic-mix" />
          </div>
          <div className="lg:hidden ">
            <Sheet>
              <SheetTrigger>
                <MenuIcon />
              </SheetTrigger>
              <SheetContent side="left">
                <TestServicesNavigator appMap={pagemap} appName="magic-mix" />
              </SheetContent>
            </Sheet>
          </div>
          <div className=" grow"></div>
          <div className="bg-slate-400 container p-1 border rounded-2xl lg:w-3/4">
            <div className="bg-white p-10 dark:bg-black border rounded-2xl  ">
              {children}
            </div>
          </div>
          <div className=" grow"></div>
          <div className="hidden lg:block"></div>
        </div>
      </div>
    </div>
  );
}
