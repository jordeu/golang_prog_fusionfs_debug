import java.util.stream.IntStream


process VCFANNO {
    tag "${meta.id}"

    container 'docker.io/scwatts/vcfanno:0.3.5'

    input:
    tuple val(meta), path(vcf)
    path annotations_dir
    each n

    output:
    path '*general.vcf.gz'

    script:
    vcf_name_stem = vcf.name.replace('.vcf.gz', '')

    """
    set -o pipefail

    vcfanno -p 4 -base-path \$(pwd) ${annotations_dir}/vcfanno_annotations.toml ${vcf} | \\
        bcftools view -o ${vcf_name_stem}.${n}.general.vcf.gz && \\
        bcftools index -t ${vcf_name_stem}.${n}.general.vcf.gz
    """
}

workflow {
    ch_inputs = Channel.fromPath(params.input)
        .splitCsv(header: true)
        .map { meta -> [meta, meta.vcf] }

    task_ns = IntStream.rangeClosed(1, 50).boxed().toList()
    VCFANNO(
        ch_inputs,
        params.annotations_dir,
        task_ns,
    )
}
