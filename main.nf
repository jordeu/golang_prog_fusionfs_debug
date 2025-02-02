import java.util.stream.IntStream


process FUSIONFS_DEBUG {
    tag "${n}"

    container 'docker.io/jordeu/fusionfs_debug:2310.01'

    input:
    path file_a
    path file_b
    each n

    script:
    """
    fusionfs_debug ${file_a} ${file_b}
    """
}

workflow {

    task_ns = IntStream.rangeClosed(1, 50).boxed().toList()

    FUSIONFS_DEBUG(
        params.file_a,
        params.file_b,
        task_ns,
    )

}
