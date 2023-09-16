import java.util.stream.IntStream


process FUSIONFS_DEBUG {
    tag "${meta.id}"

    container 'docker.io/scwatts/fusionfs_debug:2309.12'

    input:
    tuple val(meta), path(vcf)
    path annotations_dir
    each n

    script:
    """
    set -o pipefail
    fusionfs_debug ${annotations_dir}/vcfanno_annotations.toml ${vcf}
    echo completed
    """
}

workflow {
    ch_inputs = Channel.fromPath(params.input)
        .splitCsv(header: true)
        .map { meta -> [meta, meta.vcf] }

    task_ns = IntStream.rangeClosed(1, 50).boxed().toList()
    FUSIONFS_DEBUG(
        ch_inputs,
        params.annotations_dir,
        task_ns,
    )
}
