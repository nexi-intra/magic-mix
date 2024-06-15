"use client";
import { cn } from "@/app/koksmat/utils";
import { AppMap } from "../services/magic-people";
import { usePathname } from "next/navigation";
import Link from "next/link";

export function TestServicesNavigator(props: {
  appMap: AppMap;
  appName: string;
}) {
  const { appMap, appName } = props;
  const pathname = usePathname();
  return (
    <div className="mr-4">
      {appMap.services
        .sort((a, b) => {
          return a.name.localeCompare(b.name);
        })
        .map((service) => {
          const href = `/magic/services/${appName}/${service.name}`;
          return (
            <div key={service.name}>
              <Link
                className={cn(
                  "hover:underline",
                  pathname === href ? "font-bold" : ""
                )}
                href={href}
              >
                {service.name}
              </Link>
            </div>
          );
        })}
    </div>
  );
}
