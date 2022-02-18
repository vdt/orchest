"""Wraps complex k8s api calls to keep orchest/_core more readable.

This module is especially useful for orchest-ctl due to the use of
async_req=True in the k8s python SDK that happens here.
"""

from typing import List, Optional

from kubernetes import client as k8s_client

from app import config
from app.connections import k8s_apps_api


def get_orchest_deployments(
    deployments: Optional[List[str]] = None,
) -> List[Optional[k8s_client.V1Deployment]]:
    if deployments is None:
        deployments = config.ORCHEST_DEPLOYMENTS
    threads = []
    for name in deployments:
        t = k8s_apps_api.read_namespaced_deployment(
            name, config.ORCHEST_NAMESPACE, async_req=True
        )
        threads.append(t)

    responses = []
    for t in threads:
        try:
            deployment = t.get()
            responses.append(deployment)
        except k8s_client.ApiException as e:
            if e.status == 404:
                responses.append(None)
            else:
                raise
    return responses
