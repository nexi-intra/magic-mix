import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Sheet,
  SheetClose,
  SheetContent,
  SheetDescription,
  SheetFooter,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";

export function TestServicesCall(props: {
  name: string;
  children: React.ReactNode;
}) {
  const { children, name } = props;
  return (
    <Sheet>
      <SheetTrigger asChild>
        <Button>Test {name}</Button>
      </SheetTrigger>
      <SheetContent side="bottom" className="h-[90%]">
        <SheetHeader>
          <SheetTitle>Edit profile</SheetTitle>
          <SheetDescription>
            Make changes to your profile here. Click save when you're done.
          </SheetDescription>
        </SheetHeader>
        {children}
      </SheetContent>
    </Sheet>
  );
}
