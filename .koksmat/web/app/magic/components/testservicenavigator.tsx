"use client";
import { cn } from "@/koksmat/utils";
import { AppMap } from "../services/magic-mix";
import { usePathname } from "next/navigation";

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
          return (
            <div key={service.name}>
              <div className="text-xl my-2">{service.name}</div>
              <ul className="ml-2">
                {service.endpoints
                  .sort((a, b) => {
                    return a.name.localeCompare(b.name);
                  })
                  .map((endpoint) => {
                    const href = `/magic/services/${appName}/${service.name}/${endpoint.name}`;
                    return (
                      <li key={endpoint.name}>
                        <a
                          className={cn(
                            "hover:underline",
                            pathname === href ? "font-bold" : ""
                          )}
                          href={href}
                        >
                          {endpoint.name}
                        </a>
                      </li>
                    );
                  })}
              </ul>
            </div>
          );
        })}
    </div>
  );
}
