// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Rke.Inputs
{

    public sealed class ClusterCloudProviderAwsCloudConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (list maxitems:1)
        /// </summary>
        [Input("global")]
        public Input<Inputs.ClusterCloudProviderAwsCloudConfigGlobalArgs>? Global { get; set; }

        [Input("serviceOverrides")]
        private InputList<Inputs.ClusterCloudProviderAwsCloudConfigServiceOverrideArgs>? _serviceOverrides;

        /// <summary>
        /// (list)
        /// </summary>
        public InputList<Inputs.ClusterCloudProviderAwsCloudConfigServiceOverrideArgs> ServiceOverrides
        {
            get => _serviceOverrides ?? (_serviceOverrides = new InputList<Inputs.ClusterCloudProviderAwsCloudConfigServiceOverrideArgs>());
            set => _serviceOverrides = value;
        }

        public ClusterCloudProviderAwsCloudConfigArgs()
        {
        }
    }
}