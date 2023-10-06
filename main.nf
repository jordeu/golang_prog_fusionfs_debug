plugins {
    id 'nf-amazon'
}

docker {
    enabled = true
}

fusion {
    enabled = true
}

wave {
    enabled = true
}

aws {
    batch {
        jobRole = 'arn:aws:iam::843407916570:role/NextflowApplicationDevSta-TaskBatchInstanceRolesas-19GNOYR22W9AP'
        volumes = '/mnt/local_ephemeral/:/tmp/'
    }
}

process {
    executor = 'awsbatch'
    scratch = false

    errorStrategy = 'ignore'

    withName: FUSIONFS_DEBUG {
        queue = 'nextflow-task-4cpu_32gb'
        cpus = 4
        memory = 30.GB

        publishDir = [
            path: { "${params.outdir}" },
            mode: 'copy'
        ]
    }
}
