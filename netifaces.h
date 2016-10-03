//
//  netifaces.h
//  netifaces
//
//  Created by Sung-Taek, Kim on 10/2/16.
//  Copyright © 2016 PocketCluster. All rights reserved.
//

#ifndef __NETIFACES_H__
#define __NETIFACES_H__

#ifndef WIN32

#include <stdbool.h>

typedef struct Gateway {
    struct Gateway*    next;
    unsigned char      family;
    bool               is_default;
    char*              ifname;
    char*              addr;
} Gateway;

// find all system gateways & return linked list results
int find_system_gateways(Gateway** results);

// release the results
void release_gateways(Gateway** results);

// find the first ip4 default gateway in the results
Gateway* find_default_ip4_gw(Gateway** results);

// find the first ip6 default gateway in the results
Gateway* find_default_ip6_gw(Gateway** results);

#endif

#endif /* netifaces_h */
